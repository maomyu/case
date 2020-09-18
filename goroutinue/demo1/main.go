package main

import (
	"fmt"
	"sync"
	"time"
)
var (
	wg sync.WaitGroup
)
func main() {
	var Ball int
	table := make(chan int)
	wg.Add(1)
	go player(1,table)
	wg.Add(1)
	go player(2,table)
	wg.Add(1)
	go player(3,table)

	table <- Ball

	wg.Wait()
}

func player(id int,table chan int) {
	for {
		ball := <-table
		fmt.Println("玩家",id,"接球：",ball,"并开始击球：",ball+1)
		ball++
		time.Sleep(1 * time.Second)
		table <- ball
	}
	wg.Done()
}