proto-gen:
	protoc --go_out=chatserver --go_opt=paths=source_relative --go-grpc_out=chatserver --go-grpc_opt=paths=source_relative chat.proto
