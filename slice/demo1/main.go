package main

import (
	"fmt"
)

func main() {
	var array =[10]int{1,2,3,4,5,6,7,8,9,10}

	var slice = array[6:7]

	fmt.Println(slice)
	fmt.Println("lenth of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[6])
}

/*
[7]
lenth of slice:  1
capacity of slice:  4
true
*/