package pool

import (
	"log"
	"sync"
)

// Worker handles all the work
type Worker struct {
	Id			int
	TaskChan	chan *Task
	QuitChan	chan bool
}

// NewWorker returns new instance of worker
func NewWorker(taskChan chan *Task, id int) *Worker {
	return &Worker{
		Id:			id,
		TaskChan:	taskChan,
		QuitChan:	make(chan bool),
	}
}

// Start starts the worker
func (w *Worker) Start(wg *sync.WaitGroup) {
	log.Printf("worker %d starting\n", w.Id)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for task := range w.TaskChan {
			Process(w.Id, task)
		}
	}()
}

// StartBackground starts the worker in background waiting
func (w *Worker) StartBackground() {
	log.Printf("starting worker %d in background\n", w.Id)

	for {
		select {
		case task := <-w.TaskChan:
			Process(w.Id, task)
		case <-w.QuitChan:
			return
		}
	}
}

// Stop quits the worker
func (w *Worker) Stop() {
	log.Printf("closing worker %d\n", w.Id)

	go func() {
		w.QuitChan <- true
	}()
}