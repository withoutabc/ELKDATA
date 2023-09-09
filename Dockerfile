FROM alpine:latest

WORKDIR /app

COPY ./data/dynamic/cmd/main .
COPY ./IP2LOCATION-LITE-DB11.BIN .
COPY ./html/visit.html ./html/slow.html ./html/

EXPOSE 5888

ENV TZ=Asia/Shanghai

CMD ["./main"]