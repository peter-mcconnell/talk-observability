FROM golang:1.10-stretch

RUN mkdir -p /root/go/src/github.com/pemcconnell/talk-observability
WORKDIR /root/go/src/github.com/pemcconnell/talk-observability
ENV GOPATH=/root/go
ENV PATH="${GOPATH}/bin:${PATH}"

# use watcher so we can edit and see changes without restarts
RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

# demo app
COPY main.go ./main.go
RUN go get ./...

CMD ["watcher"]
