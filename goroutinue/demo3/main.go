package main

import (
	"fmt"
	"time"
)

func producer(seq int,ch chan int,d time.Duration){
	var i int
	for {
		ch <-i
		//fmt.Println(seq,"号工人","生产了",i)
		i++
		time.Sleep(d)
	}
}

func consume(seq int,out chan int,d time.Duration){
	for x := range out {
		fmt.Println(seq,"消费者消费了",x)
		time.Sleep(d)
	}
}
// 多协程写入
func main() {
	ch := make(chan int)
	out := make(chan int)
	for i :=0;i<2;i++{

	}
	go producer(1,ch, 100*time.Millisecond)
	go producer(2,ch, 250*time.Millisecond)
	go consume(3,out,150*time.Millisecond)


	for i := range ch {
		out <- i
	}
}
