FROM golang:1.14

WORKDIR /go/src/app
COPY . .

# Set GOPATH Environment Variable
ENV GOPATH /go

# Change working directory
WORKDIR $GOPATH/src/server/

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
