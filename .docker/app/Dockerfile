FROM golang:1.18-alpine

RUN apk update && apk  add git
RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD ./app /go/src/app

RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.40.4

CMD ["air", "-c", ".air.toml"]