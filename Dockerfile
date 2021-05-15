FROM golang:1.13-alpine
WORKDIR /go/src/github.com/arimax04/gketest
COPY . .
RUN go build .
CMD ["gketest"]

