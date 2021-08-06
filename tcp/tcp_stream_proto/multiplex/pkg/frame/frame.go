package frame

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/panjf2000/gnet"
)

const HeaderSize = 4
const MaxLength = 100

type Frame []byte

func (f Frame) Encode(c gnet.Conn, payload []byte) ([]byte, error) {
	result := make([]byte, 0)
	buf := bytes.NewBuffer(result)
	length := uint32(HeaderSize + len(payload))

	if err := binary.Write(buf, binary.BigEndian, length); err != nil {
		s := fmt.Sprintf("packet length error: %v", err)

		return nil, errors.New(s)
	}

	n, err := buf.Write(payload)
	if err != nil {
		s := fmt.Sprintf("packet frame payload error: %v", err)

		return nil, errors.New(s)
	}

	if n != len(payload) {
		s := fmt.Sprintf("packet frame payload length error: %v", err)

		return nil, errors.New(s)
	}

	return buf.Bytes(), nil
}

func (f Frame) Decode(c gnet.Conn) ([]byte, error) {
	//预读取，检查Frame完整性
	var frameLength uint32

	if n, header := c.ReadN(HeaderSize); n == HeaderSize {
		buf := bytes.NewBuffer(header)
		_ = binary.Read(buf, binary.BigEndian, &frameLength)

		if frameLength > MaxLength {
			c.ResetBuffer()

			return nil, errors.New("length too large")
		}

		if n, frame := c.ReadN(int(frameLength)); n == int(frameLength) {
			c.ShiftN(int(frameLength))

			return frame[HeaderSize:], nil
		} else {
			return nil, errors.New("not enough frame payload")
		}
	}

	return nil, errors.New("not enough frame length")
}