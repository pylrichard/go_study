package pool

import (
	"log"
	"sync"
	"time"
)

const (
	taskSize = 128
)

type Pool struct {
	Tasks 		[]*Task
	Concurrency int
	Collector	chan *Task
	WaitGroup	sync.WaitGroup
}

// NewPool init a new pool with the given tasks and concurrency
func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:	tasks,
		Concurrency: concurrency,
		Collector: make(chan *Task, taskSize),
	}
}

// Run start all worker within the pool and blocks until it's finished
func (p *Pool) Run() {
	start := time.Now()
	for i := 0; i < p.Concurrency; i++ {
		worker := NewWorker(p.Collector, i)
		worker.Start(&p.WaitGroup)
	}

	for _, task := range p.Tasks {
		p.Collector <-task
	}
	close(p.Collector)

	p.WaitGroup.Wait()
	elapsed := time.Since(start)
	log.Printf("took %s\n", elapsed)
}