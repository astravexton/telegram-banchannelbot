FROM golang:alpine AS builder

WORKDIR /go/src/github.com/astravexton/telegram-banchannelbot
COPY . .

RUN apk update && \
    apk add --no-cache git bash && \
    go get -d -v ./... && \
    go install

FROM alpine:latest

COPY --from=builder /go/bin/telegram-banchannelbot /usr/local/bin/telegram-banchannelbot

CMD ["telegram-banchannelbot"]
