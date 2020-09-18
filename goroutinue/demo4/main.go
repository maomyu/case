package main

import (
	"fmt"
	"sync"
	"time"
)


type Factory struct{
	tasksCh chan int
}

//多协程消费

func worker(seq int,tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			fmt.Println("当前没有任务")
			return
		}
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Println(seq," worker finished processing task", task)
	}
}
func (f *Factory) Run(wg *sync.WaitGroup, workers int) {
	for i := 0; i < workers; i++ {
		go worker(i,f.tasksCh, wg)
	}
	for i := 0; i < 100; i++ {
		f.tasksCh <- i
	}
	close(f.tasksCh)
}

func main() {
	var wg sync.WaitGroup
	f :=&Factory{
		tasksCh :make(chan int),
	}
	wg.Add(36)
	go f.Run(&wg,36)
	wg.Wait()
}
