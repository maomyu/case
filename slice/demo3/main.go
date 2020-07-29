package main

import (
	"fmt"
)

//order[low:high:max]操作意思是对order进行切片，新切片范围是[low, high),新切片容量是max
func main() {
	orderLen := 4
	order := make([]uint16, 2 * orderLen)

	order = []uint16{1,2,3,4,5,6,7,8,9,10}

	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
}