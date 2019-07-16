module github.com/kinsprite/product-test

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	google.golang.org/grpc v1.22.0
)

replace google.golang.org/grpc v1.22.0 => github.com/grpc/grpc-go v1.22.0
