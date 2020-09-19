package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	testHttpIdleConnection()

}

//测试空闲链接不可用
func testHttpIdleConnection() {

	var myHttpTransport = &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     time.Second * 60,
	}

	client := &http.Client{Transport: myHttpTransport}
	count := 0
	for {
		count++
		resp, err := client.Get("http://118.24.111.175/")

		if err != nil {
			fmt.Println(count, "error_1", err)
			continue

		}
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(count, "error_2", err)
			continue
		}

		resp.Body.Close()

		fmt.Println(count, string(body), err)

		time.Sleep(time.Second * 120)
	}

}
