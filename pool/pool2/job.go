package pool2

import "log"

type Job interface {
	Start(worker *Worker) error
}

type PrintJob struct {
	Index int
}

func (j *PrintJob) Start(worker *Worker) error {
	log.Printf("job %s - %d\n", worker.Name, j.Index)

	return nil
}