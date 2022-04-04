proto-gen:
	protoc --go_out=services --go_opt=paths=source_relative --go-grpc_out=services --go-grpc_opt=paths=source_relative chat.proto
