FROM golang:1.13.15 AS builder

WORKDIR /build

COPY ./main.go .

RUN CGO_ENABLED=0 go build -o fortune-api


FROM alpine

RUN apk add fortune

COPY --from=builder /build/fortune-api /opt/fortune-api

RUN chmod +x /opt/fortune-api

ENTRYPOINT ["/opt/fortune-api"]
