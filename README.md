# 🧮 gRPC Calculator Service

Một ví dụ demo service **Calculator** sử dụng **Go + gRPC + Protocol Buffers**.

## 📂 Cấu trúc thư mục

calculator/
├── calculatorrpc/ # File .proto và code sinh ra
│ ├── calculator.proto
│ ├── calculator.pb.go
│ └── calculator_grpc.pb.go
├── server/ # gRPC server
│ └── main.go
├── client/ # gRPC client
│ └── main.go
├── Makefile # Script build/run
└── README.md # Tài liệu project


## 🚀 Cài đặt

### 1. Yêu cầu
- Go >= 1.20
- [Protocol Buffers Compiler (`protoc`)](https://github.com/protocolbuffers/protobuf/releases)
- Plugin Go cho protobuf:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --version

make proto
make run-server
make run-client
Add Result: 15
Subtract Result: 5
make build
