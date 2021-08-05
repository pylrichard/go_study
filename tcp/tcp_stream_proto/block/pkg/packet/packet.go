package packet

import (
	"bytes"
	"fmt"
	"github.com/bytedance/gopkg/lang/mcache"
)

/*
	Packet定义：PacketHeader + PacketBody
	PacketHeader：cmdId 1byte
	PacketBody：id 8bytes
				如果是req 任意字节 payload
				如果是ack 1byte result
 */

const (
	//CmdConn 0x01，连接请求包
	CmdConn = 0x01 + iota
	//CmdSubmit 0x02，消息发送请求包
	CmdSubmit
)

const (
	//CmdConnAck 0x81，连接请求的响应包
	CmdConnAck = 0x80 + iota
	//CmdSubmitAck 0x82，消息发送请求的响应包
	CmdSubmitAck
)

const IdSize = 8
const CmdIdSize = 1

type Packet interface {
	//Decode []byte -> struct
	Decode([]byte) error
	//Encode struct -> []byte
	Encode() ([]byte, error)
}

type Submit struct {
	Id      string
	Payload []byte
}

func (s *Submit) Decode(body []byte) error {
	s.Id = string(body[:IdSize])
	s.Payload = body[IdSize:]

	return nil
}

func (s *Submit) Encode() ([]byte, error) {
	return bytes.Join([][]byte{
						[]byte(s.Id[:IdSize]),
						s.Payload}, nil), nil
}

type SubmitAck struct {
	Id 		string
	Result	uint8
}

func (sa *SubmitAck) Decode(body []byte) error {
	sa.Id = string(body[:IdSize])
	sa.Result = body[IdSize]

	return nil
}

func (sa *SubmitAck) Encode() ([]byte, error) {
	return bytes.Join([][]byte{
						[]byte(sa.Id[:IdSize]),
						{sa.Result}}, nil), nil
}

func Decode(packet []byte) (Packet, error) {
	defer mcache.Free(packet)
	cmdId := packet[0]
	body := packet[CmdIdSize:]

	switch cmdId {
	case CmdConn:
		return nil, nil
	case CmdConnAck:
		return nil, nil
	case CmdSubmit:
		s := Submit{}
		err := s.Decode(body)
		if err != nil {
			return nil, err
		}

		return &s, nil
	case CmdSubmitAck:
		sa := SubmitAck{}
		err := sa.Decode(body)
		if err != nil {
			return nil, err
		}

		return &sa, nil
	default:
		return nil, fmt.Errorf("unknown cmdId [%d]", cmdId)
	}
}

func Encode(p Packet) ([]byte, error) {
	var cmdId uint8
	var body []byte
	var err error

	switch t := p.(type) {
	case *Submit:
		cmdId = CmdSubmit
		body, err = p.Encode()
		if err != nil {
			return nil, err
		}
	case *SubmitAck:
		cmdId = CmdSubmitAck
		body, err = p.Encode()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown type [%s]", t)
	}

	return bytes.Join([][]byte{
						{cmdId}, body}, nil), nil
}