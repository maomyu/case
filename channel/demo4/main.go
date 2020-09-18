package main

import "fmt"

func main() {
	A := make(chan bool, 1)
	B := make(chan bool)
	Exit := make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			if ok := <-A; ok {
				fmt.Println("A = ", 2*i-1)
				B <- true
			}
		}
	}()
	go func() {
		defer func() {
			close(Exit)
		}()
		for i := 1; i <= 10; i++ {
			if ok := <-B; ok {
				fmt.Println("B : ", 2*i)
				A <- true
			}
		}
	}()

	A <- true
	<-Exit
}
