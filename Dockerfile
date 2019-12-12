FROM golang:1.13-alpine

RUN apk update && apk add --no-cache build-base

WORKDIR /app

COPY . .

RUN go build -o bin

CMD ["/app/bin"]
