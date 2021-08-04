package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"go/go_study/tcp/tcp_stream_proto/block/pkg/frame"
	"go/go_study/tcp/tcp_stream_proto/block/pkg/packet"

	"github.com/lucasepe/codename"
)

func main() {
	var wg sync.WaitGroup
	num := 6

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			startClient()
		}()
	}
	wg.Wait()
}

func startClient() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println("dial error: ", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")

	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}

	var counter int
	codec := frame.NewCodec()
	go func() {
		for {
			//read from connection
			//handle ack
			ackPayload, err := codec.Decode(conn)
			if err != nil {
				panic(err)
			}

			pkt, err := packet.Decode(ackPayload)
			sa, ok := pkt.(*packet.SubmitAck)
			if !ok {
				panic("not submit ack")
			}
			log.Printf("the result of submit ack[%s] is %d\n", sa.Id, sa.Result)
		}
	}()

	for {
		//send submit
		counter++
		id := fmt.Sprintf("%08d", counter)
		//generate random pronounceable code name
		//example:
		//hopeful-toad-men-133b
		//blessed-man-thing-2bdc
		payload := codename.Generate(rng, 4)
		s := &packet.Submit{
			Id:		 id,
			Payload: []byte(payload),
		}

		framePayload, err := packet.Encode(s)
		if err != nil {
			panic(err)
		}
		log.Printf("send submit id = %s, payload = %s, frame length = %d\n",
				s.Id, s.Payload, len(framePayload) + frame.HeaderSize)

		err = codec.Encode(conn, framePayload)
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}