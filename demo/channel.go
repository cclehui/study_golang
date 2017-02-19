package main;

import(
    "fmt"
    "os"
    "time"
    "runtime"
)

func testCount(i int, channel chan int) {
    fmt.Printf("i=%d\n", i);
    channel <- i;
    
}

func main() {
    fmt.Println("xxxxxx");
    fmt.Println(runtime.NumCPU());
    os.Exit(0);

    tempChannel := make(chan int, 1);
    for {
        select {
            case tempChannel <- 0:

            case tempChannel <- 1:

        }

        value := <-tempChannel;
        fmt.Println(value);
        time.Sleep(2 * time.Second);
    }
    os.Exit(1);
    channelSice := make([]chan int, 10);
    var i int;

    for i = 0; i <10; i++ {
        channelSice[i] = make(chan int, 1);
        go testCount(i, channelSice[i]);
    }

    for _, temp := range(channelSice) {
        i = <-temp
        fmt.Printf("out i = %d\n", i);

    }

    fmt.Println("end...");


}
