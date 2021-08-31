package worker

import (
	"go/go_study/pool/pool3/basic"
	"go/go_study/pool/pool3/model"
	"log"
	"sync"
	"time"
)

const (
	workerSize = 64
)

func PooledWork(allData []model.SimpleData) {
	start := time.Now()
	var wg sync.WaitGroup
	dataChan := make(chan model.SimpleData, dataSize)

	for i := 0; i < workerSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for data := range dataChan {
				basic.Process(data)
			}
		}()
	}
	for _, data := range allData {
		dataChan <-data
	}
	close(dataChan)
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("took %s\n", elapsed)
}