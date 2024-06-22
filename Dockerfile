FROM golang:1.22-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN apk update && \
    apk add --no-cache \
        libpng-dev \
        libjpeg-turbo-dev \
        libwebp-dev
RUN go build -o main cmd/main.go

CMD ["/app/main"]