FROM golang:1.23

WORKDIR /

COPY ./ ./

RUN go mod download
RUN go build -o server ./cmd/main/main.go

EXPOSE 50051

CMD ["./server"]