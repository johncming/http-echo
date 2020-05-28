FROM golang:1.14.2 AS builder

WORKDIR /go/src/http-echo

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

FROM busybox:latest

COPY --from=builder /go/src/http-echo /bin/

ENTRYPOINT ["/bin/http-echo"]
