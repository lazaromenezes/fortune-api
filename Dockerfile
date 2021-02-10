FROM alpine

RUN apk add fortune

COPY ./main /opt/fortune-api
RUN chmod +x /opt/fortune-api

EXPOSE 80

ENTRYPOINT ["/opt/fortune-api"]
