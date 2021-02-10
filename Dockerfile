FROM alpine

RUN apk add fortune

COPY ./main /opt/fortune-api
RUN chmod +x /opt/fortune-api

EXPOSE 8765

ENTRYPOINT ["/opt/fortune-api"]
