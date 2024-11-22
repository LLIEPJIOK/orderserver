FROM golang:1.23-alpine

WORKDIR /

COPY . .

RUN go mod download
RUN go build -o server ./cmd/main/main.go

EXPOSE 50051
EXPOSE 8080

CMD ["./server"]