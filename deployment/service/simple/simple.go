package simple

import (
	"github.com/kardianos/service"
	"log"
)

//service日志对象
var logger service.Logger

//实现service.Interface的Start和Stop接口
type program struct {}

func (p *program) Start(s service.Service) error {
	// Start() should not block. Do the actual work async
	go p.run()

	return nil
}

func (p *program) run() {
	log.Printf("program running...")
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func Simple() {
	svcConfig := &service.Config{
		Name:	"SimpleGoService",
		DisplayName: "simple",
		Description: "This is a example go service",
	}

	p := &program{}
	s, err := service.New(p, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		_ = logger.Error(err)
	}
}