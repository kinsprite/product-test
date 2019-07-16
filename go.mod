module github.com/kinsprite/producttest

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	go.elastic.co/apm v1.4.0
	go.elastic.co/apm/module/apmgrpc v1.4.0
	go.elastic.co/apm/module/apmhttp v1.4.0
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.22.0
)

replace google.golang.org/grpc v1.22.0 => github.com/grpc/grpc-go v1.22.0
