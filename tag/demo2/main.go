package main

import (
	"fmt"
	"reflect"
)

type A struct{
	Age int64 `json:"age"`
	Name string `json:"name"`
}

type I interface{}
func main(){
	a :=&A{1,"hah"}

	rv :=reflect.ValueOf(a)
	rt :=reflect.TypeOf(a)
	fmt.Println(rv)
	fmt.Println(rt)

	var i I
	i = a
	sv :=reflect.ValueOf(i)
	st :=reflect.TypeOf(i)
	fmt.Println(sv.String())
	fmt.Println(st.String())
	fmt.Println("99999999999999999")
	fmt.Println(true&&false)
	
}
