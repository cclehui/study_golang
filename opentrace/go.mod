module github.com/cclehui/study_golang/opentrace

go 1.13

require (
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/uber/jaeger-client-go v0.0.0-00010101000000-000000000000
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/atomic v1.5.1 // indirect
)

replace github.com/uber/jaeger-client-go => /home/godev/go_workspace/src/github.com/uber/jaeger-client-go
