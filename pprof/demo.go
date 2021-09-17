package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// 相关资料文档
// https://juejin.cn/post/7006983753905995790


var datas []string

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}

func main() {
	go func() {
		for{
			log.Printf("len: %d", Add("go-programming-tour-book"))
			time.Sleep(time.Millisecond * 500)
		}
	}()
	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}
