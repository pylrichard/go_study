package pool

import "log"

type Task struct {
	Err		error
	Data	interface{}
	Execute func(interface{}) error
}

func NewTask(execute func(interface{}) error, data interface{}) *Task {
	return &Task{Execute: execute, Data: data}
}

func Process(workerId int, task *Task) {
	log.Printf("worker %d processes task %v\n", workerId, task.Data)
	task.Err = task.Execute(task.Data)
}