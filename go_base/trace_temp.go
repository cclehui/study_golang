package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

//trace 数据查看方法 go tool trace trace.out

func mockSendToServer(url string) {
	fmt.Printf("server url: %s\n", url)
}

func main() {
	f, err := os.Create("trace.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	urls := []string{"0.0.0.0:5000", "0.0.0.0:6000", "0.0.0.0:7000"}
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mockSendToServer(url)
		}()
	}
	wg.Wait()
}
