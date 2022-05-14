FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./
RUN go build -o dns-notify

FROM alpine:3.15 AS run

ENV DNS=8.8.8.8

COPY --from=build ./app/dns-notify ./

CMD [ "./dns-notify" ]
