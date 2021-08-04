package monitor

import (
	"expvar"
	"time"
)

var (
	SubmitInTotal	*expvar.Int
	SubmitInRate	*expvar.Int
)

func init() {
	SubmitInTotal = expvar.NewInt("SubmitInTotal")
	SubmitInRate = expvar.NewInt("SubmitInRate")

	go func() {
		var lastTotal int64

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <- ticker.C:
				newTotal := SubmitInTotal.Value()
				//两秒处理的消息量之差作为处理速度
				SubmitInRate.Set(newTotal - lastTotal)
				lastTotal = newTotal
			}
		}
	}()
}