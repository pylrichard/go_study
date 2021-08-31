package worker

import (
	"log"
	"sync"
	"time"

	"go/go_study/pool/pool3/basic"
	"go/go_study/pool/pool3/model"
)

const (
	dataSize = 128
)

func NotPooledWork(allData []model.SimpleData) {
	start := time.Now()
	var wg sync.WaitGroup
	dataChan := make(chan model.SimpleData, dataSize)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range dataChan {
			wg.Add(1)
			go func(data model.SimpleData) {
				defer wg.Done()
				basic.Process(data)
			}(data)
		}
	}()
	for _, data := range allData {
		dataChan <-data
	}
	close(dataChan)
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("took %s\n", elapsed)
}