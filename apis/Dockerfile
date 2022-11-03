FROM golang:1.19.1-alpine as go-builder
WORKDIR /api
COPY . .

RUN apk add build-base &&\
    go build &&\
    chmod 777 apis

FROM alpine:latest
RUN adduser -D -h /api -u 1000 -k /dev/null api
WORKDIR /api

COPY --from=go-builder --chown=nobody:nobody /api/apis apis

EXPOSE 8090
USER api
CMD [ "./apis" ]