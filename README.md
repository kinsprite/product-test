# gRPC
```
protoc -I pb/ --go_out=plugins=grpc:pb ./pb/helloworld.proto
```

# Build

```
go mod tidy
go mod vendor
go build -mod=vendor
```
