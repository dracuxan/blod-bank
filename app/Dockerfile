FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o server .

EXPOSE 5099

CMD ["./server"]
