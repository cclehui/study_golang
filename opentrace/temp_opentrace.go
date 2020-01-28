package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/cclehui/study_golang/opentrace/tracer"
	opentracing "github.com/opentracing/opentracing-go"
)

//启动 jaeger server端服务
//./jaeger-all-in-one --collector.zipkin.http-port=9411

func main() {
	fmt.Println("aaaaaaaaaaaaaaa")
	var (
		err error
		io  io.Closer
	)
	//创建tracer对象
	tracer.Tracer, io, err = tracer.NewTracer("study_golang_opentrace_test", "127.0.0.1:6831")
	if err != nil {
		log.Fatal("tracer.NewTracer error(%v)", err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer.Tracer)

	fmt.Printf("11111, %v\n", tracer.Tracer)

	//server
	http.HandleFunc("/trace_test", traceTest)
	log.Printf("Starting server on port %d", 8002)
	err = http.ListenAndServe(fmt.Sprintf(":%d", 8002), http.DefaultServeMux)
	if err != nil {
		log.Fatalf("Cannot start server: %s", err)
	}
}

func traceTest(w http.ResponseWriter, r *http.Request) {
	log.Println("Received trace test request")

	//client
	rootSpan := tracer.Tracer.StartSpan("traceTest")
	rootSpan.SetTag("action", "traceTest")
	defer rootSpan.Finish()

	wait := sync.WaitGroup{}
	startTime := time.Now()

	for i := 1; i < 4; i++ {
		wait.Add(1)

		go func(i int) {

			spanName := fmt.Sprintf("child_span_%d", i)

			childSpan := tracer.Tracer.StartSpan(spanName, opentracing.ChildOf(rootSpan.Context()))
			time.Sleep(time.Second * time.Duration(i))

			if i > 2 {
				childSpan.LogKV("overtime", fmt.Sprintf("时间超过阈值:%d", i))
			}

			childSpan.Finish()
			wait.Done()

		}(i)

	}
	wait.Wait()
	//ctx := opentracing.ContextWithSpan(context.Background(), span)

	costTime := time.Now().Unix() - startTime.Unix()

	//log.Printf("Received result: %s\n", body)
	io.WriteString(w, fmt.Sprintf("cost time: %d", costTime))
}

/*
func onError(span opentracing.Span, err error) {
	span.SetTag(string(ext.Error), true)
	span.LogKV("key1", err)
	log.Fatal("client(%v)", err)
}
*/
