package main

import (
	"bufio"
	"fmt"
	"go/go_study/tcp/tcp_stream_proto/block/pkg/monitor"
	"log"
	"net"
	"net/http"

	"go/go_study/tcp/tcp_stream_proto/block/pkg/frame"
	"go/go_study/tcp/tcp_stream_proto/block/pkg/packet"
)

func main() {
	//monitor http server
	go func() {
		log.Fatal(http.ListenAndServe(":8090", nil))
	}()

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("listen error: ", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept error: ", err)
			break
		}
		//start a goroutine to handle the new connection
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	codec := frame.NewCodec()
	buf := bufio.NewReader(c)

	for {
		//read from the connection
		//decode the frame to get the payload
		//the payload is a packet which is not decoded
		framePayload, err := codec.Decode(buf)
		if err != nil {
			log.Println("handleConn frame decode error: ", err)
			return
		}
		//handle the packet
		ackFramePayload, err := handlePacket(framePayload)
		if err != nil {
			log.Println("handleConn handle packet error: ", err)
			return
		}
		//write ack frame to the connection
		err = codec.Encode(c, ackFramePayload)
		if err != nil {
			log.Println("handleConn: frame encode error: ", err)
			return
		}
		//每处理完一个请求，计数器+1
		monitor.SubmitInTotal.Add(1)
	}
}

func handlePacket(payload []byte) ([]byte, error) {
	p, err := packet.Decode(payload)
	if err != nil {
		log.Println("handlePacket packet decode error: ", err)
		return nil, err
	}

	switch p.(type) {
	case *packet.Submit:
		s := p.(*packet.Submit)
		log.Printf("recv submit: id = %s, payload = %s\n", s.Id, string(s.Payload))
		sa := &packet.SubmitAck{
			Id:		s.Id,
			Result: 0,
		}
		ackFramePayload, err := packet.Encode(sa)
		if err != nil {
			log.Println("handlePacket packet encode error: ", err)
			return nil, err
		}

		return ackFramePayload, nil
	default:
		return nil, fmt.Errorf("unknown packet type")
	}
}