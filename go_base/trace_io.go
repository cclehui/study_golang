package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/trace"
	"strconv"
	"sync"
)

func main() {

	f, _ := os.Create("trace.log")
	trace.Start(f)     //开启trace 捕获并输出到 f所指向的 trace.log文件中
	defer trace.Stop() //结束trace
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			a := 0
			for i := 0; i < 1e6; i++ {
				a += 1
				if i == 1e6/2 {
					bytes, _ := ioutil.ReadFile(`test.txt`)
					inc, _ := strconv.Atoi(string(bytes))
					a += inc
				}
			}
			wg.Done()
			fmt.Println(a)
		}()
	}

	wg.Wait()

}
