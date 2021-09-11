package pool3

import (
	"log"
	"math/rand"
	"os"
	"time"

	"go/go_study/pool/pool3/basic"
	"go/go_study/pool/pool3/model"
	"go/go_study/pool/pool3/pool"
	"go/go_study/pool/pool3/worker"

	"github.com/urfave/cli"
)

// UsePool https://github.com/Joker666/goworkerpool.git
func UsePool() {
	// Prepare the data
	var allData []model.SimpleData
	for i := 0; i < 128; i++ {
		data := model.SimpleData{ Id: i }
		allData = append(allData, data)
	}

	// Prepare the task
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

	// Use cli package to make command line tool
	// to explore various examples we have built here
	app := &cli.App{
		Name: "worker_pool",
		Usage: "check different work loads with worker pool",
		Action: func(c *cli.Context) error {
			log.Println("you need more parameters")

			return nil
		},
		Commands: []cli.Command{
			{
				Name: "basic",
				Usage: "Run synchronously",
				Action: func(c *cli.Context) error {
					basic.Work(allData)

					return nil
				},
			},
			{
				Name:  "not_pooled",
				Usage: "Run without any pooling",
				Action: func(c *cli.Context) error {
					worker.NotPooledWork(allData)

					return nil
				},
			},
			{
				Name:  "pooled",
				Usage: "Run with pooling",
				Action: func(c *cli.Context) error {
					worker.PooledWork(allData)

					return nil
				},
			},
			{
				Name:  "pooled_error",
				Usage: "Run with pooling that handles errors",
				Action: func(c *cli.Context) error {
					worker.PooledWorkError(allData)

					return nil
				},
			},
			{
				Name:  "pool",
				Usage: "Run robust worker pool",
				Action: func(c *cli.Context) error {
					p.Run()

					return nil
				},
			},
			{
				Name:  "pool_bg",
				Usage: "Run robust worker pool in background",
				Action: func(c *cli.Context) error {
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

					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}