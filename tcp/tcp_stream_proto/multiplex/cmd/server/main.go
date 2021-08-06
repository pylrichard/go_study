package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"go/go_study/tcp/tcp_stream_proto/multiplex/pkg/frame"
	"go/go_study/tcp/tcp_stream_proto/multiplex/pkg/packet"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
)

type CodecServer struct {
	*gnet.EventServer
	addr		string
	isMultiCore bool
	isAsync		bool
	codec		gnet.ICodec
	workerPool *goroutine.Pool
}

/*
	go run main.go --port 8888 --isMultiCore=true
 */
func main() {
	var port int
	var isMultiCore bool

	flag.IntVar(&port, "port", 8888, "server port")
	flag.BoolVar(&isMultiCore, "isMultiCore", true, "isMultiCore")
	flag.Parse()

	addr := fmt.Sprintf("tcp://:%d", port)
	runCodecServer(addr, isMultiCore, false, nil)
}

func runCodecServer(addr string, isMultiCore, isAsync bool, codec gnet.ICodec) {
	var err error
	codec = frame.Frame{}
	srv := &CodecServer{
		addr: addr, isMultiCore: isMultiCore,
		isAsync: isAsync, codec: codec,
		workerPool: goroutine.Default(),
	}
	err = gnet.Serve(srv, addr, gnet.WithMulticore(isMultiCore),
		gnet.WithTCPKeepAlive(6 * time.Minute), gnet.WithCodec(codec),
	)
	if err != nil {
		panic(err)
	}
}

func (svr * CodecServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("codec server is listening on %s (multicore: %t, loop: %d)\n",
				srv.Addr.String(), srv.Multicore, srv.NumEventLoop)

	return
}

func (svr *CodecServer) React(framePayload []byte, c gnet.Conn) (ackFramePayload []byte, action gnet.Action) {
	var p packet.Packet

	p, err := packet.Decode(framePayload)
	if err != nil {
		log.Println("react packet decode error: ", err)
		//close connection
		action = gnet.Close
		return
	}

	switch p.(type) {
	case *packet.Submit:
		s := p.(*packet.Submit)
		log.Printf("recv submit: id = %s, payload = %s\n", s.Id, string(s.Payload))
		sa := &packet.SubmitAck{
			Id:	s.Id,
			Result: 0,
		}
		ackFramePayload, err = packet.Encode(sa)
		if err != nil {
			log.Println("react packet encode error: ", err)
			action = gnet.Close
			return
		}

		return
	default:
		return nil, gnet.Close
	}
}