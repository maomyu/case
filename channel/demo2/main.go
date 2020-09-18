package main

import (
	"time"
	"math/rand"
	"sync"
	"log"
)

//多个生产者，单个消费者
func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	// ...
	const MaxRandomNumber = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	// ...
	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})
	// stopCh is an 停止通道的信号.

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				// 这个select目的是结束该goroutinue、

				select {
				case <- stopCh:
					return
				default:
				}

				select {
				case <- stopCh:
					return
				case dataCh <- rand.Intn(MaxRandomNumber):
				}
			}
		}()
	}

	// 消费者
	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == MaxRandomNumber-1 {
				//////////////////////////////
				close(stopCh)
				return
			}

			log.Println(value)
		}
	}()

	// ...
	wgReceivers.Wait()
}

// channel 在没有任何 goroutine 引用的时候会自行关闭，而不需要显示进行关闭。
// 关闭通道，会广播到所有的协程；