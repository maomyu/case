package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// 通用的http请求
func httpDo(urls string,method string, values map[string][]string) string{
	client := &http.Client{}

	urlmap := url.Values{}
	urlmap = values

	parms := ioutil.NopCloser(strings.NewReader(urlmap.Encode())) //把form数据编下码
	req, err := http.NewRequest(method, urls, parms)
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println(err)
		return ""
	}
	time.Sleep(3*time.Second)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}
func PraseURL(str string)string{
	data :=strings.Split(str,"\n")

	return data[2]
}
func main(){
	var m sync.WaitGroup
	for i:=0;i<100;i++{
		m.Add(1)
		go func(){
			str :=httpDo("http://172.16.110.1:1260/live/test.m3u8?domain=hdl.fake.fake","GET",map[string][]string{})
			streamulr :=PraseURL(str)
			httpDo(streamulr,"GET", map[string][]string{})
			fmt.Println(streamulr)
			m.Done()
			time.Sleep(2*time.Second)
		}()

	}
	m.Wait()
	fmt.Println("end.........")
}