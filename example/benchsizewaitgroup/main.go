package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"git2.qingtingfm.com/car/car-common/sizedwaitgroup"
	render "git2.qingtingfm.com/neirong/library/base/logrender"
	"git2.qingtingfm.com/neirong/library/goroutine"
)

var swg = sizedwaitgroup.New(10)

func main() {
	fmt.Println("111111111")
	runtime.GOMAXPROCS(2)
	benchSizeWaitGroup()

	//select {}
	time.Sleep(time.Second * 2)
	fmt.Println("done,xxxxxxxxx")
	//swg.Wait()
	select {}
	fmt.Println("done,yyyyyyyyyy")
}

func benchSizeWaitGroup() {
	conf := &goroutine.Config{
		Config: &render.Config{
			Stdout: false,
			OutDir: "",
		},
	}
	goroutine.Init(conf)

	for i := 0; i < 100; i++ {
		go getOneDataBySW()
	}

}

func getOneDataBySW() {
	goGroup := goroutine.New("channelsProgramInfos", goroutine.SetMaxWorker(10, true))
	ctx := context.Background()
	for i := 0; i < 1000; i++ {
		goName := fmt.Sprintf("getOneDataBySW_%d", i)
		//go func(i int) {
		goGroup.Go(ctx, goName, func(ctx context.Context) error {
			swg.Add()        // 获取资源
			defer swg.Done() // 释放资源

			for j := 0; j <= 65; j++ {
				httpQuery()
			}

			return nil
		})
	}

	goGroup.Wait()
}

// 获取http 数据 50 ms
func httpQuery() {
	//fmt.Println(i)
	time.Sleep(time.Millisecond * 50)
}
