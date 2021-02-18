# grpc-example
 
A gRPC implementation using Go lang

## 🛠 Installing

You need to have Go lang installed on your machine.

### Compiling protobuffers and gRPC

```bash
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```

## 🏃‍♂️ Running

First of all, you need to run the server:

```bash
go run cmd/server/server.go
```

After that, you just need to run the client simultaneously:

```bash
go run cmd/client/client.go
```

