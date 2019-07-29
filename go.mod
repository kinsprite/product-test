module github.com/kinsprite/producttest

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.10
	github.com/json-iterator/go v1.1.6
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/segmentio/kafka-go v0.3.0
	go.elastic.co/apm v1.4.0
	go.elastic.co/apm/module/apmgrpc v1.4.0
	go.elastic.co/apm/module/apmhttp v1.4.0
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3
	google.golang.org/grpc v1.22.0
)

replace google.golang.org/grpc v1.22.0 => github.com/grpc/grpc-go v1.22.0
