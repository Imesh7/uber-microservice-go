genererate grpc files

protoc --go_out=gen/driver_auth --go_opt=paths=source_relative --go-grpc_out=gen/driver_auth --go-grpc_opt=paths=source_relative --proto_path=api api/driver_auth.proto