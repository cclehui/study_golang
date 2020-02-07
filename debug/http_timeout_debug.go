package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	count := 0

	for {
		rand.Seed(time.Now().Unix())
		count++

		for i := 1; i < 5; i++ {

			size := 50
			start := 50000000 + rand.Intn(40000000)
			limit := 50

			//target := "http://pic01.babytreeimg.com/img/event/icon1x1.gif"
			target := "http%3A%2F%2Fpic01.babytreeimg.com%2Fimg%2Fevent%2Ficon1x1.gif"

			//url := fmt.Sprintf("http://go.babytree.com/go_discuss/api/http_timeout/debug?size=%d&start=%d&limit=%d&url=%s&http_type=raw", size, start, limit, target)
			//url := fmt.Sprintf("http://172.16.9.216:8897/go_discuss/api/http_timeout/debug?size=%d&start=%d&limit=%d&url=%s&http_type=raw", size, start, limit, target)
			url := fmt.Sprintf("http://go.chenlehui.babytree-dev.com/go_discuss/api/http_timeout/debug?size=%d&start=%d&limit=%d&url=%s&http_type=raw", size, start, limit, target)
			//url := fmt.Sprintf("http://go.babytree-pre.com/go_discuss/api/http_timeout/debug?size=%d&start=%d&limit=%d&url=%s&http_type=raw", size, start, limit, target)
			//url := fmt.Sprintf("http://go.babytree.com/go_discuss/api/http_timeout/debug?size=%d&start=%d&limit=%d&url=%s", size, start, limit, target)
			//url := fmt.Sprintf("http://go.1.babytree-test.com/go_discuss/api/http_timeout/debug?size=%d&start=%d&limit=%d", size, start, limit)
			fmt.Printf("%d, %d, %s\n", count, i, url)
			go func(url string) {
				resp, _ := http.Get(url)
				if resp == nil {
					return
				}
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Printf("response, %v\n", body)
				defer resp.Body.Close()

			}(url)
		}

		time.Sleep(time.Second * 1)
	}
}
