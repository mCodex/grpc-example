# grpc-example
 
A gRPC implementation using Go lang

## Compiling protobuffers and gRPC

```bash
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```