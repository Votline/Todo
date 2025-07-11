FROM golang:1.24.4-alpine

WORKDIR /app

COPY . .

WORKDIR /app/todo-service

RUN go mod download

RUN go build -o todo-service .

CMD ["./todo-service"]
