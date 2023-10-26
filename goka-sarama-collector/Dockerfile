FROM golang:1.19-alpine as builder
WORKDIR /install

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download

COPY ./main.go ./main.go
COPY ./event.go ./event.go
COPY ./config.go ./config.go
RUN CGO_ENABLED=0 go build -o /install/goka-sarama-collector


FROM gcr.io/distroless/static-debian11
WORKDIR /app

COPY --from=busybox:1.35.0-uclibc /bin/sh /bin/sh
COPY --from=busybox:1.35.0-uclibc /bin/sleep /bin/sleep

COPY --from=builder /install/goka-sarama-collector /app

COPY ./wait-and-start.sh ./