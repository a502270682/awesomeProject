package main

import (
	"fmt"
	"time"
)

func main() {
	paramGet()
}

func chanT() {
	ch := make(chan int)
	go func() {
		defer close(ch) // ch 需要close，不然读取时会deadlock
		for i := 1; i < 5; i++ {
			ch <- i
		}

	}()
	for num := range ch {
		fmt.Println(num)
	}
}

func selectT() {
	a := make(chan int, 1)
	// select 是随机执行的
	for i := 0; i < 3; i++ {
		select {
		case a <- 1:
		case a <- 2:
		case a <- 3:

		}
		fmt.Println(<-a)
	}
}

func paramGet() {
	ch := make(chan int)
	go fmt.Println(<-ch) // go语句后面的函数调用，其参数会先求值 造成死锁
	ch<-1
	time.Sleep(time.Second*5)

}
