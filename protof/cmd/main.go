package main

import (
	"fmt"
	"github.com/yuwe1/case/protof/listen"
	"github.com/yuwe1/case/protof/server"
	"net/http"

)

func main(){
	l ,_:=listen.RetryListen("tcp", "0.0.0.0:9999")
	srv :=server.Server{
		Server :&http.Server{},
	}
	http.HandleFunc("/latency/{delaytime}",srv.GetServerHttp)
	fmt.Println("start.....")
	srv.Server.Serve(l)
}