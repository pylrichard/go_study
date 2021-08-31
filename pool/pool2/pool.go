package pool2

import (
	"fmt"
	"log"
	"sync"
)

const (
	workerSize = 6
	queueSize = 20
	jobSize = 64
)

type Pool struct {
	Name string

	WorkerSize int
	Workers    []*Worker

	QueueSize	int
	Queue		chan Job
}

func (p *Pool) Init() {
	// maintain minimum 1 worker
	if p.WorkerSize < 1 {
		p.WorkerSize = 1
	}
	p.Workers = []*Worker{}
	for i := 0; i < p.WorkerSize; i++ {
		worker := &Worker{
			Id:		i,
			Name:	fmt.Sprintf("%s-worker-%d", p.Name, i),
		}
		p.Workers = append(p.Workers, worker)
	}
	// maintain min queue size as 1
	if p.QueueSize < 1{
		p.QueueSize = 1
	}
	p.Queue = make(chan Job, p.QueueSize)
}

func (p *Pool) Start() {
	for _, worker := range p.Workers {
		worker.Start(p.Queue)
	}
	log.Println("all workers started")
}

func (p *Pool) Stop() {
	close(p.Queue)

	var wg sync.WaitGroup
	for _, worker := range p.Workers {
		wg.Add(1)

		go func(w *Worker) {
			defer wg.Done()

			w.Stop()
		}(worker)
	}

	wg.Wait()
	log.Println("all workers stopped")
}

func UsePool() {
	pool := &Pool{
		Name:       "test",
		WorkerSize: workerSize,
		QueueSize:  queueSize,
	}
	pool.Init()
	pool.Start()
	defer pool.Stop()

	for i := 1; i < jobSize; i++ {
		job := &PrintJob{
			Index: i,
		}
		pool.Queue <-job
	}
}