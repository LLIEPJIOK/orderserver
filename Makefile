run:
	go run cmd/main/main.go

build:
	go build cmd/main/main.go

generate-proto:
	protoc -I ./proto \
	--go_out ./pkg/api --go_opt paths=source_relative \
	--go-grpc_out ./pkg/api --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./pkg/api --grpc-gateway_opt paths=source_relative \
	./proto/order/order.proto