package main

import (
	"fmt"
	"time"
)

//并发竞争  竞态检查

func main() {

	data := make([]int, 10000)

	for i := 0; i < 10000; i++ {
		tempIndex := i
		go func() {
			data[tempIndex] = tempIndex

		}()
	}

	select {
	case <-time.After(time.Second * 3):
		fmt.Println("111111111")
	}
}
