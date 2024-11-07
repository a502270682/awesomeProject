package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type GoroutinePool struct {
	tasks   chan Task
	wg      sync.WaitGroup
	maxSize int
}

func NewGoroutinePool(maxSize int) *GoroutinePool {
	pool := &GoroutinePool{
		tasks:   make(chan Task),
		maxSize: maxSize,
	}

	for i := 0; i < maxSize; i++ {
		workerID := i
		go pool.worker(workerID)
	}

	return pool
}

func (p *GoroutinePool) worker(workerId int) {
	for task := range p.tasks {
		task()
		fmt.Println("worker", workerId, "done")
		p.wg.Done()
	}
}

func (p *GoroutinePool) AddTask(task Task) {
	p.wg.Add(1)
	p.tasks <- task
}

func (p *GoroutinePool) Wait() {
	p.wg.Wait()
	close(p.tasks)
}

func main() {
	pool := NewGoroutinePool(3) // 最大并发数为3
	fmt.Println(time.Now())
	for i := 0; i < 10; i++ {
		i := i // 为了避免闭包问题
		pool.AddTask(func() {
			defer fmt.Printf("Task %d done\n", i)
			time.Sleep(2 * time.Second) // 模拟某个耗时操作
		})
	}

	pool.Wait() // 等待所有任务完成
	fmt.Println("All tasks completed.", time.Now())
}
