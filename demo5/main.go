// go语言运行shell命令
package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sync"
)

//ffmpeg -re -i ~/Desktop/ffmpeg-test/testlong.mp4 -codec copy -f flv rtmp://172.16.100.1:1966/live/test3
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func (){
		var cmd *exec.Cmd
		cmd = exec.Command("ffmpeg","-re -i ~/Desktop/ffmpeg-test/testlong.mp4 -codec copy -f flv", "\"rtmp://10.200.20.28:1966/jztest4/test?domain=pili-publish.jztest4.cloudvdn.com\"")
		//显示运行的命令
		fmt.Println(cmd.Args)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		if err :=cmd.Start();err!=nil{
			fmt.Println(err)
		}

		reader := bufio.NewReader(stdout)
		//实时循环读取输出流中的一行内容
		for {
			line, err2 := reader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Println(line)
		}
		//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
		cmd.Wait()
		wg.Done()
	}()
	wg.Wait()
}