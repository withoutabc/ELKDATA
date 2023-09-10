FROM alpine:latest

WORKDIR /app

COPY ./data/dynamic/cmd/main .
COPY ./IP2LOCATION-LITE-DB11.BIN .
COPY ./html/visit.html ./html/slow.html ./html/

EXPOSE 5888

RUN apk update && apk add tzdata

CMD ["./main"]