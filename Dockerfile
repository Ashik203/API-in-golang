FROM golang:1.22.0-alpine

WORKDIR /app

COPY . .

WORKDIR /app/cmd/server

RUN go get -d -v ./...

RUN go build -o /app/app .

EXPOSE 8090

CMD ["/app/app"]
