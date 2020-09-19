package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
void output(char *str) {
	    usleep(1000000);
		    printf("%s\n", str);
		}
*/
import "C"
import (
	"log"
	"net/http"
	"runtime"
	"sync"
	"unsafe"

	_ "net/http/pprof"
)

//http://172.16.9.216:9090/debug/pprof/
//参考资料: https://mp.weixin.qq.com/s?__biz=MjM5MDUwNTQwMQ==&mid=2257484187&idx=1&sn=5277c9c5d8a0ed5bd78f9216c9a60c1a&chksm=a539194d924e905b9bffbfc74b0c25e5d14045f902605f7a8abd4e7a352eb72565469c4ef558&mpshare=1&scene=1&srcid=0919KfuN2qBlfAFfW4dTIQUu&sharer_sharetime=1600481561277&sharer_shareid=51b6c5cc6082300df9f5098a0e391717&key=ca85356d1be876fdc615967ca3e265be2c67f54763a5fc146e23473a3117b38c96a252bc7245e3779d6588a4aa8d62f4c485a1f29c600064401a9bb794d52311199c6c5be32acf53f2de37878b7fc35531c53e5b2a7f9e3197025bc6272c12609c730d4948255e5c69b5b1b2e23bde2859321b68821ec9231ce0f17efce48aac&ascene=1&uin=MjEyMzI5MDcwMw%3D%3D&devicetype=Windows+10+x64&version=62090538&lang=zh_CN&exportkey=AYXHmUjDyCexmnUghyN4DqE%3D&pass_ticket=27tbfThORKN0LSlAHgvsz9QtyBtGFxtRC3ZiF1L5hVKEkVHidV3E6EuoaTD3XvHS&wx_header=0

func init() {
	go http.ListenAndServe(":9090", nil)
}

//测试 线程数量不会缩容
func main() {
	for i := 0; i < 100; i++ {
		go func() {
			str := "hello cgo"
			//change to char*
			cstr := C.CString(str)
			C.output(cstr)
			C.free(unsafe.Pointer(cstr))

		}()
	}

	//启动释放线程的server http api
	killThreadService()

	select {}
}

//释放线程
func sayhello(wr http.ResponseWriter, r *http.Request) {
	KillOne()
}

func killThreadService() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":10003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// KillOne kills a thread
func KillOne() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		runtime.LockOSThread()
		return
	}()

	wg.Wait()
}
