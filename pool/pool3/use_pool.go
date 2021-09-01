package pool3

import (
	"go/go_study/pool/pool3/pool"
	"log"
	"time"
)

func UsePool() {
	var allTask []*pool.Task
	for i := 0; i < 64; i++ {
		task := pool.NewTask(func(data interface{}) error {
			taskId := data.(int)
			time.Sleep(100 * time.Millisecond)
			log.Printf("task %d processed", taskId)

			return nil
		}, i)
		allTask = append(allTask, task)
	}

	p := pool.NewPool(allTask, 6)
	p.Run()
}