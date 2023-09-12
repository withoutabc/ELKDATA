FROM alpine:latest

WORKDIR /app

COPY ./data/dynamic/cmd/main .
COPY ./front_end/visit.html ./front_end/slow.html ./front_end/visit.css ./front_end/visit.js ./html/

EXPOSE 5888

RUN apk update && apk add tzdata

CMD ["./main"]