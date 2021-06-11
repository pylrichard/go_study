package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func workForever() {
	for {
		go counter()
		time.Sleep(1 * time.Second)
	}
}

func httpGet(w http.ResponseWriter, r *http.Request) {
	betterCounter()
	//counter()
}

func WorkForever() {
	go workForever()
	http.HandleFunc("/get", httpGet)
	_ = http.ListenAndServe(":8080", nil)
}