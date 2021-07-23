package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("11111111111")
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)

	for i := 0; i < 30; i++ {
		go func(x int) {
			cond.L.Lock()
			fmt.Println(x, " 获取锁")
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println(x, " 被唤醒")
			time.Sleep(time.Second * 5)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Signal...")
	cond.Signal()
	/*
		time.Sleep(time.Second)
		cond.Signal()
		time.Sleep(time.Second * 3)
		cond.Broadcast()
	*/

	fmt.Println("Broadcast...")
	time.Sleep(time.Minute)
}

// kkkkkkkkkkkkkk
type SizedSyncCond struct {
	Size int

	cond *sync.Cond
}

func NewSizedSyncCond(size int) *SizedSyncCond {
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)
	return &SizedSyncCond{
		Size: size,
		cond: cond,
	}
}

func (ssc *SizedSyncCond) LockAndWait() {
	ssc.cond.L.Lock()
	ssc.cond.Wait()
}

// func (ssc *SizedSyncCond) Singal() {
func (ssc *SizedSyncCond) Done() {
	ssc.cond.Singal()
	ssc.cond.L.Unlock()
}
