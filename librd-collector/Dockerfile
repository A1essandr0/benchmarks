FROM golang:1.20.2-bullseye AS builder
WORKDIR /install

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download

COPY ./event.go ./event.go
COPY ./config.go ./config.go
COPY ./main.go ./main.go
RUN go build -o /install/librd-collector


FROM gcr.io/distroless/base-debian11
WORKDIR /app

COPY --from=busybox:1.35.0-uclibc /bin/sh /bin/sh
COPY --from=busybox:1.35.0-uclibc /bin/sleep /bin/sleep

COPY --from=builder /install/librd-collector /app

COPY ./wait-and-start.sh ./
