package main

import (
	"fmt"
	"net"
)

func main() {
	//mc := memcache.New("172.16.9.128:11211")
	network := "tcp"
	address := "118.24.111.175:11211"
	dialer := net.Dialer{}
	conn, err := dialer.Dial(network, address)
	fmt.Println(conn, err)
	fmt.Printf("222, %#v\n", dialer)
	/*
		count := 0
			mc := memcache.New("118.24.111.175:11211")
			for {
				count++
				err := mc.Set(&memcache.Item{Key: "cclehui_foo", Value: []byte("my value")})

				//it, err := mc.Get("cclehui_foo")
				//fmt.Println(count, it, "error:", err)
				fmt.Println(count, "error:", err)

				time.Sleep(time.Millisecond * 1000)
			}
	*/

	select {}
}
