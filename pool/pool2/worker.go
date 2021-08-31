package pool2

import "log"

type Worker struct {
	Id			int
	Name		string
	StopChan 	chan bool
}

func (w *Worker) Start(jobQueue chan Job) {
	w.StopChan = make(chan bool)
	successChan := make(chan bool)

	go func() {
		successChan <-true
		for {
			// take job from queue
			job := <-jobQueue
			if job != nil {
				_ = job.Start(w)
			} else {
				log.Printf("worker %s to be stopped\n", w.Name)
				w.StopChan <-true
				break
			}
		}
	}()

	// wait for the worker to start
	<-successChan
}

func (w *Worker) Stop() {
	// wait for the worker to stop, blocking
	_ = <-w.StopChan
	log.Printf("worker %s stopped\n", w.Name)
}