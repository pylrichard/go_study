package pool

import "log"

// Task encapsulates a work item that should go in a work
type Task struct {
	Err		error
	Data	interface{}
	Execute func(interface{}) error
}

// NewTask init a new task based on a given work function
func NewTask(execute func(interface{}) error, data interface{}) *Task {
	return &Task{Execute: execute, Data: data}
}

// Process handles the execution of task
func Process(workerId int, task *Task) {
	log.Printf("worker %d processes task %v\n", workerId, task.Data)
	task.Err = task.Execute(task.Data)
}