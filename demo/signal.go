package main

import (
    "fmt"
    "os"
    "syscall"
    "os/signal"
)

func main() {

    var sig_chan = make(chan os.Signal)

    signal.Notify(sig_chan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
    //signal.Notify(sig_chan, syscall.SIGHUP)

    sig := <-sig_chan
    fmt.Println(sig)

    fmt.Println("server closed by " + sig.String())


}
