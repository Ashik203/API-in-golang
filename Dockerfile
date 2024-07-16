FROM golang:1.16.3-alpine3.13

WORKDIR /app

COPY . .

WORKDIR /app/cmd/server

RUN go get -d -v ./...

RUN go build -o /app/app .

EXPOSE 8090

CMD ["/app/app"]
