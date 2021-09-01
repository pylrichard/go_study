package pool

import (
	"log"
	"sync"
	"time"
)

const (
	taskSize = 128
)

// Pool is the worker pool
type Pool struct {
	Tasks 		[]*Task
	Workers 	[]*Worker

	Concurrency 		int
	Collector			chan *Task
	RunningBackground 	chan bool
	WaitGroup			sync.WaitGroup
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

// AddTask adds a task to the pool
func (p *Pool) AddTask(task *Task) {
	p.Collector <-task
}

// RunBackground runs the pool in background
func (p *Pool) RunBackground() {
	go func() {
		for {
			log.Println("waiting for task to come in...")
			time.Sleep(6 * time.Second)
		}
	}()

	for i := 1; i <= p.Concurrency; i++ {
		worker := NewWorker(p.Collector, i)
		p.Workers = append(p.Workers, worker)
		go worker.StartBackground()
	}

	for _, task := range p.Tasks {
		p.Collector <-task
	}

	p.RunningBackground = make(chan bool)
	<-p.RunningBackground
}

// Stop stops background workers
func (p *Pool) Stop() {
	for _, worker := range p.Workers {
		worker.Stop()
	}
	p.RunningBackground <-true
}