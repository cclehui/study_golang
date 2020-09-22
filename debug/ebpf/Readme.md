## Go Dynamic Tracepoint example.
## golang ebpf 使用demo
This is an example of using [gobpf](https://github.com/iovisor/gobpf) to trace arguments of the function for our example [application](https://github.com/pixie-labs/pixie/blob/main/demos/simple-gotracing/app.go).

### Dependencies
This requires [libbcc](https://github.com/iovisor/bcc/blob/master/INSTALL.md) to be installed.

### Build
```
go build -o app ebpf_app.go

go build -o trace ebpf_trace.go
```

### Run
```
./trace --binary ./app  # To trace the example app.


curl http://127.0.0.1:9090/e

```

### 参考资料
https://blog.pixielabs.ai/blog/ebpf-function-tracing/post/
https://github.com/pixie-labs/pixie/blob/main/demos/simple-gotracing/trace_example/trace.go
