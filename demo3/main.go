package main

import (
	"fmt"
	"time"
)

type Bar struct{
	percent int64 //百分比float32
	cur int64 //当前的进度位置
	total int64 //总进度
	rate string //进度条
	graph string  //显示符号
}

//start 和 total支持断点续传
func (bar *Bar)NewOption(start,total int64){
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph ="█"
	}
	bar.percent = bar.getPercent()
	//i += 2的步长，因为百分比总是从0到100，而进度条长度最长为50个字符，这也就意味着，每增长2%，进度条就要涨一格，因此，这里的步长为2
	for i:=0;i<int(bar.percent);i+=1{
		bar.rate +=bar.graph
	}
}

func (bar *Bar) NewOptionWithGraph(start, total int64, graph string) {
	bar.graph = graph
	bar.NewOption(start, total)
}

//当前进度/总量=已完成百分比
func (bar *Bar) getPercent() int64 {
	return int64(float32(bar.cur)/float32(bar.total)*100)
}

//进度条显示
func (bar *Bar)Play(cur int64){
	bar.cur = cur
	last :=bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last  && bar.percent%2 ==0{
		bar.rate +=bar.graph
	}
	fmt.Printf("\r%s%3d%%  %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) Finish(){
	fmt.Println()
}

func main() {
	var bar Bar
	bar.NewOption(30, 100)
	for i:= 30; i<=100; i++{
		time.Sleep(100*time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
}
