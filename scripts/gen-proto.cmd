protoc --proto_path=protos/user-service --gofast_out=plugins=grpc:. user.proto
protoc --proto_path=protos/user-service --proto_path=protos/user-service --gofast_out=plugins=grpc:. user.proto