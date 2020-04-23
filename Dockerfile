FROM node:12-alpine

WORKDIR /app

COPY . .

RUN npm i
RUN npm run build

FROM node:12-alpine

WORKDIR /app

COPY ./themes/default .

RUN npm i
RUN npm build

FROM golang:1.13-alpine

RUN apk update && apk add --no-cache build-base

WORKDIR /app

COPY . .

RUN go build -o bin

FROM alpine:3.11

WORKDIR /app

COPY ./.env.example ./.env

COPY --from=0 /app/public /app/public

COPY --from=1 /app /app/themes/default

COPY --from=2 /app/bin /app/bin

EXPOSE 3000

CMD ["/app/bin"]
