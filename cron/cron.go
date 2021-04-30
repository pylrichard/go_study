package cron

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"
	"time"

	"github.com/robfig/cron/v3"
)

type greetingJob struct {
	Name string
}

func (g greetingJob) Run() {
	fmt.Println("Hello ", g.Name)
}

type panicJob struct {
	count int
}

func (p *panicJob) Run() {
	p.count++
	if p.count == 1 {
		panic("ops!")
	}

	fmt.Println("hello world")
}

type delayJob struct {
	count int
}

func (d *delayJob) Run() {
	//增加了一个2s延迟，输出中间隔变为2s，而不是定时的1s
	time.Sleep(2 * time.Second)
	d.count++
	log.Printf("%d: hello world\n", d.count)
}

type skipJob struct {
	count int32
}

func (s *skipJob) Run() {
	atomic.AddInt32(&s.count, 1)
	log.Printf("%d: hello world\n", s.count)
	if atomic.LoadInt32(&s.count) == 1 {
		time.Sleep(2 * time.Second)
	}
}

func TimeRule() {
	c := cron.New()
	/*
		添加定时任务
		参数1以字符串形式指定触发时间规则
		参数2是一个无参函数，每次触发时调用
		@every 1s表示每秒触发一次，@every后加一个时间间隔，表示每隔多长时间触发一次
		@every 1m2s表示每隔1分2秒触发一次
		time.ParseDuration()支持的格式都可以用在这里
	 */
	c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 sec")
	})
	c.AddFunc("@hourly", func() {
		fmt.Println("every hour")
	})
	c.AddFunc("30 * * * *", func() {
		fmt.Println("every hour on the half hour")
	})
	/*
		启动一个新goroutine做定时循环检测，最后加了一行time.Sleep()防止主goroutine退出
	 */
	c.Start()
	for {
		time.Sleep(time.Second)
	}
}

func TimeZone() {
	loc, _ := time.LoadLocation("Asia/Beijing")
	c := cron.New(cron.WithLocation(loc))
	c.AddFunc("0 6 * * ?", func() {
		fmt.Println("Every 6 o'clock at Beijing")
	})
	c.AddFunc("CRON_TZ=Asia/Tokyo 0 6 * * ?", func() {
		fmt.Println("Every 6 o'clock at Tokyo")
	})
	c.Start()
	for {
		time.Sleep(time.Second)
	}
}

func Job() {
	c := cron.New()
	//使用自定义结构可以让任务携带状态(Name字段)
	c.AddJob("@every 1s", greetingJob{"pyl"})
	c.Start()
	for {
		time.Sleep(time.Second)
	}
}

func Parser() {
	//创建Parser对象，以位格式传入使用域，默认时间格式使用5个域
	//parser := cron.NewParser(
	//	cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	//)
	//c := cron.New(cron.WithParser(parser))
	c := cron.New(cron.WithSeconds())
	c.AddFunc("1 * * * * *", func () {
		fmt.Println("every 1 second")
	})
	c.Start()
	for {
		time.Sleep(time.Second)
	}
}

func Logger() {
	c := cron.New(cron.WithLogger(
		//cron.VerbosePrintfLogger()包装log.Logger，logger会详细记录cron内部的调度过程
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags)),
	))
	c.AddFunc("@every 1s", func() {
		fmt.Println("hello world")
	})
	c.Start()
	for {
		time.Sleep(time.Second)
	}
}

func Recover() {
	c := cron.New()
	//panicJob在第一次触发时，触发了panic。因为有cron.Recover()保护，后续任务还能执行
	c.AddJob("@every 1s",
		cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&panicJob{}))
	c.Start()

	for {
		time.Sleep(time.Second)
	}
}

func DelayIfStillRunning() {
	c := cron.New()
	c.AddJob("@every 1s",
		cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&delayJob{}))
	c.Start()

	for {
		time.Sleep(time.Second)
	}
}

func SkipIfStillRunning() {
	c := cron.New()
	//注意观察时间，第1个与第2个输出之间相差3s，因为跳过了2次执行
	c.AddJob("@every 1s",
		cron.NewChain(cron.SkipIfStillRunning(cron.DefaultLogger)).Then(&skipJob{}))
	c.Start()

	for {
		time.Sleep(time.Second)
	}
}