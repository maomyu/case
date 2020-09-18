package main

import (
	"fmt"
	"sync"
)

type info struct {
	sync.RWMutex
	data int
}

func rwMutex(count int)  {
	c := make(chan struct{},count * 3)
	l := info{data:0}
	go func(){
		for i := 0; i < count; i++ {
			go func() {
				l.RLock()
				d := l.data
				fmt.Printf("我读取到了data，值为:%d\n",d)
				l.RUnlock()
				c <- struct{}{}
			}()
		}
	}()

	go func(){
		for i := 0; i < count; i++ {
			go func(i int) {
				l.Lock()
				l.data += i
				fmt.Printf("我把data的值加了%d变成了%d\n",i,l.data)
				l.Unlock()
				c <- struct{}{}
			}(i)
		}
	}()

	go func(){
		for i := 0; i < count; i++ {
			go func(i int) {
				l.Lock()
				l.data -= i
				fmt.Printf("我把data的值减了%d变成了%d\n",i,l.data)
				l.Unlock()
				c <- struct{}{}
			}(i)
		}
	}()

	for i := 0; i < count * 3; i++ {
		<-c
	}
	fmt.Printf("data的最终结果应该为0，实际结果为：%d",l.data)
}

func rwWithoutMutex(count int)  {
	c := make(chan struct{},count * 3)
	l := 0
	for i := 0; i < count; i++ {
		go func() {
			fmt.Printf("我读取到了data，值为:%d\n",l)
			c <- struct{}{}
		}()
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			l += i
			fmt.Printf("我把data的值加了%d变成了%d\n",i,l)
			c <- struct{}{}
		}(i)
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			l -= i
			fmt.Printf("我把data的值减了%d变成了%d\n",i,l)
			c <- struct{}{}
		}(i)
	}
	for i := 0; i < count * 3; i++ {
		<-c
	}
	fmt.Printf("不安全读写时data的最终结果应该为0，实际结果为：%d",l)
}

func main(){
	rwMutex(10)
}