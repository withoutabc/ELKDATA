FROM alpine:latest

WORKDIR /app

COPY ./data/dynamic/cmd/main .
COPY ./html/visit.html ./html/slow.html ./html/

EXPOSE 5888

RUN apk update && apk add tzdata

CMD ["./main"]