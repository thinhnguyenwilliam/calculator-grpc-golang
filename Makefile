#Sinh code từ proto:make proto
#Chạy server:make run-server
#Chạy client:make run-client
#Build binary:make build
#Build lại từ đầu:make all

PROTO_DIR=calculatorrpc
PROTO_FILES=$(PROTO_DIR)/calculator.proto

# Module name theo go.mod
MODULE=github.com/thinhnguyen-com/calculator

# Tạo code từ .proto
proto:
	protoc --go_out=. --go-grpc_out=. $(PROTO_FILES)

# Chạy server
run-server:
	go run ./server

# Chạy client
run-client:
	go run ./client

# Format code
fmt:
	go fmt ./...

# Build server & client binary
build:
	go build -o bin/server ./server
	go build -o bin/client ./client

# Clean binary
clean:
	rm -rf bin

# Regenerate + build lại
all: clean proto build
