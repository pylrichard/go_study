package pool

import (
	"log"
	"sync"
)

type Worker struct {
	Id			int
	TaskChan	chan *Task
}

func NewWorker(taskChan chan *Task, id int) *Worker {
	return &Worker{
		Id:			id,
		TaskChan:	taskChan,
	}
}

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