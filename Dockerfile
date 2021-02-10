FROM alpine

RUN apk add fortune

COPY ./main /opt/fortune-api
RUN chmod +x /opt/fortune-api

ENTRYPOINT ["/opt/fortune-api"]
