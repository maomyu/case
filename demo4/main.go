package main

type I interface{}
type A struct {

}
type B struct {
	A
}

func (a *A)t1(){

}
func (a *A)t2(){

}
func (b *B)t3(){

}
func main(){
	b :=&B{}
	b.t1()
}