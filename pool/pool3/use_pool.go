package pool3

import (
	"go/go_study/pool/pool3/basic"
	"go/go_study/pool/pool3/model"
	"log"
)

func UsePool() {
	var allData []model.SimpleData
	for i := 0; i < 128; i++ {
		data := model.SimpleData{ Id: i }
		allData = append(allData, data)
	}
	log.Printf("start processing all work\n")
	basic.Work(allData)
}