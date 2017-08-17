FROM golang

COPY . /go/src/github.com/my-repo/pubsub
WORKDIR /go/src/github.com/my-repo/pubsub

RUN go get ./...
RUN go build
CMD pubsub
