package basic

import (
	"fmt"
	"go/go_study/pool/pool3/model"
	"log"
	"time"
)

// Work does the heavy lifting
func Work(allData []model.SimpleData) {
	start := time.Now()
	for i, _ := range allData {
		Process(allData[i])
	}
	elapsed := time.Since(start)
	log.Printf("took %s\n", elapsed)
}

// Process handles the job
func Process(data model.SimpleData) {
	log.Printf("start processing %d\n", data.Id)
	time.Sleep(100 * time.Millisecond)
	log.Printf("finish processing %d\n", data.Id)
}

// ProcessError handles the job while also handling error
func ProcessError(data model.SimpleData, errorsChan chan<-error) {
	log.Printf("start processing %d\n", data.Id)
	time.Sleep(100 * time.Millisecond)
	if data.Id % 33 == 0 {
		errorsChan <- fmt.Errorf("error on job %v", data.Id)
	} else {
		log.Printf("finish processing %d\n", data.Id)
	}
}