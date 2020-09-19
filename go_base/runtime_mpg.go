package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	testAllGs()
}

func sayhello(wr http.ResponseWriter, r *http.Request) {}

//测试全局的协程变量  allgs增长
//go tool pprof --inuse_space -svg -output heap.svg http://172.16.9.216:9090/debug/pprof/heap
func testAllGs() {
	for i := 0; i < 100000; i++ {
		go func() {
			time.Sleep(time.Second * 10)
		}()
	}
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
