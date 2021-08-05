package frame

import (
	"encoding/binary"
	"io"

	"github.com/bytedance/gopkg/lang/mcache"
)

/*
	Frame定义：FrameHeader + FramePayload(Packet)
	FrameHeader：4bytes Frame总长度
	FramePayload：Packet
 */

const HeaderSize = 4

type Payload []byte

type StreamCodec interface {
	//Encode data -> frame，并写入io.Writer
	Encode(w io.Writer, p Payload) error
	//Decode 从io.Reader中提取Frame Payload，并返回给上层
	Decode(r io.Reader) (Payload, error)
}

type Codec struct {}

func NewCodec() StreamCodec {
	return &Codec{}
}

func (c *Codec) Encode(w io.Writer, p Payload) error {
	totalLen := HeaderSize + int32(len(p))
	err := binary.Write(w, binary.BigEndian, &totalLen)
	if err != nil {
		return err
	}

	//make sure all data will be written to outbound stream
	for {
		//write the frame payload to outbound stream
		num, err := w.Write(p)
		if err != nil {
			return err
		}
		//写入完毕退出循坏
		if num >= len(p) {
			break
		}
		//未写入完毕则定位到已写入数据偏移，循环写入
		if num < len(p) {
			p = p[num:]
		}
	}

	return nil
}

func (c *Codec) Decode(r io.Reader) (Payload, error) {
	var totalLen int32
	err := binary.Read(r, binary.BigEndian, &totalLen)
	if err != nil {
		return nil, err
	}

	buf := mcache.Malloc(int(totalLen - HeaderSize))
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}