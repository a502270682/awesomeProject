package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

var cronCtl Ctl
var vnLocal *time.Location

type Ctl struct {
	C *cron.Cron
}

func init() {
	var err error
	vnLocal, err = time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		panic(errors.Wrap(err, "failed to LoadLocation"))
	}
	cronCtl.C = cron.New(cron.WithSeconds(), cron.WithLocation(vnLocal), cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
}

func main() {
	// withSeconds 才能 设置6位参数
	//c := cron.New(cron.WithSeconds(), cron.WithLogger(
	//	cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	channel := make(chan int, 0)
	jobWorker := GreetingJob{
		Name: "gfy",
	}

	// AddFunc
	//entryId, _ := c.AddFunc("@every 5s", jobWorker.Run)

	// AddJob
	entryId, err := cronCtl.C.AddJob("0 36 17 3 * *", &jobWorker) // cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&jobWorker)
	if err != nil {
		panic(err)
	}
	// 支持linux中时间规则 * * * * *（分钟，小时，天，周，月）
	fmt.Println(fmt.Sprintf("my job id is %d", entryId))
	//c.AddJob("@every 1s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&panicJob{}))
	cronCtl.C.Start()
	<-channel
}

type GreetingJob struct {
	Name  string
	count int
}

func (g *GreetingJob) Run() {
	g.count++
	//if g.count == 1 {
	//	panic("panic here")
	//} else {
	fmt.Println(fmt.Sprintf("start to do %s's job", g.Name))
	time.Sleep(1 * time.Second)
	fmt.Println(fmt.Sprintf("finish to %s's job", g.Name))
	//}

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
