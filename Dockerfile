FROM golang:1.19 as build

WORKDIR /opt/build

COPY main.go .

RUN CGO_ENABLED=0 go build main.go

FROM alpine

RUN apk add fortune

COPY --from=build /opt/build/main /opt/fortune-api

RUN chmod +x /opt/fortune-api

CMD ["/opt/fortune-api"]
