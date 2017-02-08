package main

import (
    "fmt"
    //"strconv"
    //"time"
    "sync"
)

var sum int = 0;

func test(a int, lock *sync.Mutex) {
    lock.Lock();
    sum += a;
    //fmt.Println("data:" + strconv.Itoa(a));
    fmt.Printf("i=%d, sum=%d\n", a, sum);
    lock.Unlock();
}

func main() {
    fmt.Println("xxxxxx");

    var mutexLock = &sync.Mutex{};

    for i := 1;i<10;i++ {
        go test(i, mutexLock);
    }

    for {
        mutexLock.Lock();
        if sum >=45 {
            break;
        }
        mutexLock.Unlock();
    }
    
    //time.Sleep(3 * time.Second);

}
