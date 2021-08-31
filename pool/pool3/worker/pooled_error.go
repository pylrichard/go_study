package worker

import (
	"go/go_study/pool/pool3/basic"
	"go/go_study/pool/pool3/model"
	"log"
	"sync"
	"time"
)

const (
	errorsSize = 256
)

// PooledWorkError handles tasks while also handling error
func PooledWorkError(allData []model.SimpleData) {
	start := time.Now()
	var wg sync.WaitGroup
	dataChan := make(chan model.SimpleData, dataSize)
	errorsChan := make(chan error, errorsSize)

	for i := 0; i < workerSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for data := range dataChan {
				basic.ProcessError(data, errorsChan)
			}
		}()
	}
	for _, data := range allData {
		dataChan <-data
	}
	close(dataChan)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case err := <-errorsChan:
				log.Println("finished with error:", err.Error())
			case <-time.After(1 * time.Second):
				log.Println("timeout: errors finished")
				return
			}
		}
	}()
	defer close(errorsChan)
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("took %s\n", elapsed)
}