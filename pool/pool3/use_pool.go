package pool3

import (
	"log"
	"math/rand"
	"time"

	"go/go_study/pool/pool3/pool"
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
	go func() {
		for {
			taskId := rand.Intn(64)

			if taskId % 6 == 0 {
				p.Stop()
			}

			time.Sleep(time.Duration(rand.Intn(6)) * time.Second)

			task := pool.NewTask(func(data interface{}) error {
				taskId := data.(int)
				time.Sleep(100 * time.Millisecond)
				log.Printf("task %d processed\n", taskId)

				return nil
			}, taskId)
			p.AddTask(task)
		}
	}()
	p.RunBackground()
}