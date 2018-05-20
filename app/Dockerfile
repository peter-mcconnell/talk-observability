FROM golang:1.10-stretch

RUN mkdir -p /root/go/src/github.com/pemcconnell/talk-observability
WORKDIR /root/go/src/github.com/pemcconnell/talk-observability
ENV GOPATH=/root/go

COPY main.go ./main.go
RUN go get ./...
RUN go build -o app

CMD ["./app"]
