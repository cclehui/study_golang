package main

import (
	"fmt"
	"log"

	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) React(c gnet.Conn) (out []byte, action gnet.Action) {
	out = c.Read()
	c.ResetBuffer()
	return
}

func main() {
	echo := new(echoServer)
	fmt.Println("11111111111")
	log.Fatal(gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true)))
}
