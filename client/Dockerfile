FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o client

CMD ["./client"]