FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod /go/src/github.com/aikon001/colorapiserver/
WORKDIR /go/src/github.com/aikon001/colorapiserver/
RUN go mod download
COPY . /go/src/github.com/aikon001/colorapiserver/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/aikon001 github.com/aikon001/colorapiserver/

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/aikon001/colorapiserver/build/aikon001 /usr/bin/aikon001
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/aikon001"]