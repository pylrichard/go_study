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

/*
	多goroutine场景下要注意的问题：
	1. 系统资源耗尽 2. GC STW延迟增大 3. 上下文切换延迟

	goroutine的数量和执行效率之间往往存在一个平衡点，如果控制不好，极有可能因资源争抢而出现阻塞
	或因对外部服务的高频访问而将服务拖死。应对此类问题的一种解决思路就是：goroutine pool

	设计思路：
	Go中channel是语言级支持的一种数据类型，实现了协程间基于消息传递的通信方式，是线程安全的
	先将任务数据以统一的数据结构TaskData推入到一个数据队列(channel)
	启动指定数量的工作协程，对数据队列中的任务数据进行抢占式消费
	即在所有的工作协程中，谁抢到任务数据谁就处理，处理完成后，将任务结果以统一的数据结构TaskResult写入另一个结果队列(channel)中
	然后再继续获取任务数据进行执行，如果工作协程发现数据队列为空，则工作协程退出
	这种方式不会给任何工作协程分配指定数量的任务，效率高的可以多处理，效率低的允许少处理
	从总体上达到处理时间最小化。等待所有任务数据处理完成后，将任务结果统一返回给上层调用逻辑
 */

//TaskData 任务数据
type TaskData struct {
	id   string
	guid string
	data interface{}
}

//TaskResult 任务结果
type TaskResult struct {
	id    string
	guid  string
	data  interface{}
	err   error
}

type TaskExecutor func(taskData TaskData) TaskResult

var ErrorDataListLen = errors.New("InputTaskList length must be greater than 0")

func runPool(workerNum int, dataList []TaskData, executor TaskExecutor) ([]TaskResult, error) {
	resultList := make([]TaskResult, 0)
	dataListLen := len(dataList)
	if dataListLen == 0 {
		return nil, ErrorDataListLen
	}
	dataChan := make(chan TaskData, dataListLen)
	for _, data := range dataList {
		dataChan <- data
	}
	resultChan := make(chan TaskResult, dataListLen)
	var lock sync.RWMutex
	for i := 0; i < workerNum; i++ {
		go func() {
			for {
				lock.Lock()
				if len(dataChan) == 0 {
					lock.Unlock()
					break
				}
				data := <-dataChan
				lock.Unlock()
				// 执行任务处理数据得到结果
				result := executor(data)
				resultChan <- result
			}
		}()
	}

	// 阻塞接收任务结果
	for i := 0; i < dataListLen; i++ {
		resultList = append(resultList, <-resultChan)
	}

	return resultList, nil
}

func executeTask(taskData TaskData) TaskResult {
	resultData := map[string]string{"id": taskData.id, "guid": taskData.guid}
	time.Sleep(1 * time.Second)
	log.Printf("Worker id=%s, guid=%s\n", taskData.id, taskData.guid)
	taskResult := TaskResult{
		id:   taskData.id,
		guid: taskData.guid,
		data: resultData,
		err:  nil,
	}

	return taskResult
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

	taskSize := 200
	dataList := make([]TaskData, 0)
	for i := 0; i < taskSize; i++ {
		u := uuid.NewV4()
		data := &TaskData{
			id:   strconv.Itoa(i),
			guid: u.String(),
		}
		dataList = append(dataList, *data)
	}
	result, _ := runPool(*workerNum, dataList, executeTask)
	for _, v := range result {
		log.Println("result=", v.id, v.data)
	}
}