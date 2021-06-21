package sync

import (
	"fmt"
	"sync"
	"time"
)

/*
sync.Cond 有三个方法，它们分别是：

Wait，阻塞当前协程，直到被其他协程调用 Broadcast 或者 Signal 方法唤醒，使用的时候需要加锁，使用 sync.Cond 中的锁即可，也就是 L 字段。

Signal，唤醒一个等待时间最长的协程。

Broadcast，唤醒所有等待的协程。

注意：在调用 Signal 或者 Broadcast 之前，要确保目标协程处于 Wait 阻塞状态，不然会出现死锁问题。
 */
//10个人赛跑，1个裁判发号施令

func race(){

	cond :=sync.NewCond(&sync.Mutex{})

	var wg sync.WaitGroup

	wg.Add(11)

	for i:=0;i<10; i++ {

		go func(num int) {

			defer  wg.Done()

			fmt.Println(num,"号已经就位")

			cond.L.Lock()

			cond.Wait()//等待发令枪响

			fmt.Println(num,"号开始跑……")

			cond.L.Unlock()

		}(i)

	}

	//等待所有goroutine都进入wait状态

	time.Sleep(2*time.Second)

	go func() {

		defer  wg.Done()

		fmt.Println("裁判已经就位，准备发令枪")

		fmt.Println("比赛开始，大家准备跑")

		cond.Broadcast()//发令枪响

	}()

	//防止函数提前返回退出

	wg.Wait()

}

