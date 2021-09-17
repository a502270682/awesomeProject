package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	n := make(chan int)
	s := make(chan int)

	go counter(n)
	go squarer(s, n)
	printer(s)

	/*
		in := make(chan <- int) 只可进不可出
		in <- 1

		out := make(<- chan int) 只可出不可进
		<- out
	*/

}
