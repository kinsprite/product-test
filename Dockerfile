# build
FROM golang:1.12.7-alpine3.10 as build

ENV PORT 8080
EXPOSE 8080

RUN mkdir /app
ADD . /app

ENV GOPROXY https://goproxy.io
ENV GIN_MODE release

WORKDIR  /app
RUN go mod vendor
RUN go build -mod=vendor -tags=jsoniter -o product-test .


# release
FROM alpine:3.10
RUN mkdir /app
COPY --from=build /app/product-test /app/product-test

WORKDIR  /app
CMD ["/app/product-test"]
