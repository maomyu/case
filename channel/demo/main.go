package main

import (
	"log"
	"sync"
)

func main(){
	//生产者读取数据
	//消费者打印

	var m sync.RWMutex

	const ProduceNums int= 15
	const ConsumeNums int= 1

	//所有的数据
	data :=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38}

	dataCh :=make(chan int,1)
	stopCh :=make(chan struct{})
	ToStop :=make(chan int,1)
	var wg sync.WaitGroup
	wg.Add(ConsumeNums)

	var signal int
	go func(){
		signal = <-ToStop
		close(stopCh)
	}()

	for i:=0;i<ProduceNums;i++{
		go func(i int){
			for {
				m.Lock()
				if len(data) == 0{
					m.Unlock()
					select{
					case ToStop <-i:
					default:
					}
				}else{
					select{
					case <-stopCh:
						return
					default:
					}
					select{
					case <-stopCh:
						return
					default:
						val := data[0]
						data = data[1:]

						//fmt.Println("生产者读了一个数值：",val)
						dataCh <- val
						m.Unlock()
					}
				}
			}
		}(i)
	}

	for i :=0;i<ConsumeNums;i++{
		go func(){
			defer wg.Done()
			for {
				// 收到停止的信号后结束
				select {
				case <- stopCh:
					return
				default:
				}

				// 收到停止的信号后结束
				select {
				case <- stopCh:
					return
				case value :=<-dataCh:
					log.Println(value)
				}
			}
		}()
	}
	wg.Wait()
	log.Println("stopped by", signal)
}
