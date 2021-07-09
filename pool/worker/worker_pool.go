package worker

import (
	"errors"
	"flag"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

//InputTask 单个待处理任务
type InputTask struct {
	Id			string
	Guid		string
	// 与任务关联的数据
	InputData	interface{}
}

//ResultTask 返回单个任务结果
type ResultTask struct {
	Id			string
	Guid		string
	OutputData	interface{}
	Err			error
	isRet		bool
}

type callBack func(task InputTask) (id, guid string, outputData interface{}, err error, isRet bool)

var ErrorTaskLen = errors.New("InputTaskList length must be greater than 0")

func DoFixedSizeWorks(concurrency int, inputTaskList []InputTask, cb callBack) ([]ResultTask, error) {
	resultTaskList := make([]ResultTask, 0)
	taskLen := len(inputTaskList)
	if taskLen == 0 {
		return nil, ErrorTaskLen
	}
	inputTaskChan := make(chan InputTask, taskLen)
	for _, task := range inputTaskList {
		inputTaskChan <- task
	}
	resultTaskChan := make(chan ResultTask, taskLen)
	var lock sync.RWMutex
	for i := 0; i < concurrency; i++ {
		go func() {
			for {
				lock.Lock()
				if len(inputTaskChan) == 0 {
					lock.Unlock()
					break
				}
				taskData := <-inputTaskChan
				lock.Unlock()
				// 调用任务处理回调方法
				id, guid, outputData, err, isRet := cb(taskData)
				resultTaskChan <- ResultTask{
					Id:			id,
					Guid:		guid,
					Err:		err,
					OutputData: outputData,
					isRet: 		isRet,
				}
			}
		}()
	}

	// 阻塞接收协程的执行结果
	for i := 0; i < taskLen; i++ {
		resultTaskList = append(resultTaskList, <-resultTaskChan)
	}

	return resultTaskList, nil
}

func Callback(task InputTask) (id, guid string, outputData interface{}, err error, isRet bool) {
	outputData = map[string]string{"id": task.Id, "guid": task.Guid}
	time.Sleep(10 * time.Second)
	isRet = true
	log.Printf("Worker id=%s, guid=%s\n", task.Id, task.Guid)

	return task.Id, task.Guid, outputData, nil, isRet
}

func init() {
	// 以时间为初始化种子
	rand.Seed(time.Now().UnixNano())
}

func UsePool() {
	workerNum := flag.Int("workerNum", runtime.NumCPU(), "worker num")
	flag.Parse()

	if *workerNum <= 0 {
		log.Println("workerNum must be greater than 0")
		return
	}

	jobSize := 200
	taskList := make([]InputTask, 0)
	for i := 0; i < jobSize; i++ {
		u := uuid.NewV4()
		input := &InputTask{
			Id:		strconv.Itoa(i),
			Guid:	u.String(),
		}
		taskList = append(taskList, *input)
	}
	output, _ := DoFixedSizeWorks(*workerNum, taskList, Callback)
	for _, v := range output {
		log.Println("result=", v.Id, v.OutputData)
	}
}