package main

import (
    "fmt"
    "io"
    "os"
    "log"
    "time"
    "syscall"
    "net/http"
    "os/signal"
)

func main() {
    http.HandleFunc("/", handlerDefault);

    var addr = ":8002";

    fmt.Println("server listen on " + addr)

    //启动server
    log.Fatal(http.ListenAndServe(addr, nil))

}

func handlerDefault(w http.ResponseWriter, request *http.Request) {
    io.WriteString(w, "Hello, time is " + time.Now().Format("2006-01-02 15:04:05"))
}
