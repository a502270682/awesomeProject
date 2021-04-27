package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

func main() {
	c := cron.New(cron.WithSeconds(),cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	channel := make(chan int, 0)
	jobWorker := GreetingJob{
		Name: "gfy",
	}

	// AddFunc
	//entryId, _ := c.AddFunc("@every 5s", jobWorker.Run)

	// AddJob
	entryId, _ := c.AddJob("@every 2s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&jobWorker))
	// 支持linux中时间规则 * * * * *（分钟，小时，天，周，月）
	fmt.Println(fmt.Sprintf("my job id is %d", entryId))
	//c.AddJob("@every 1s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&panicJob{}))
	c.Start()
	<- channel
}

type GreetingJob struct {
	Name string
	count int
}

func (g *GreetingJob) Run() {
	g.count++
	if g.count == 1 {
		panic("panic here")
	} else {
		fmt.Println(fmt.Sprintf("start to do %s's job", g.Name))
		time.Sleep(1*time.Second)
		fmt.Println(fmt.Sprintf("finish to %s's job", g.Name))
	}

	fmt.Println(g.count)
}

type panicJob struct {
	count int
}

func (p *panicJob) Run() {
	p.count++
	if p.count == 1 {
		panic("oooooooooooooops!!!")
	}

	fmt.Println(fmt.Sprintf("hello world, %d", p.count))
}
