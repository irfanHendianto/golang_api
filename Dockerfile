FROM golang:1.19.4-alpine

WORKDIR /app

COPY . .

RUN go build -o golang-api

EXPOSE 8080

CMD ./golang-api