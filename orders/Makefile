.PHONY: proto

proto:
	protoc --proto_path=proto \
		--go_out=proto/inventory --go_opt=paths=source_relative \
		--go-grpc_out=proto/inventory --go-grpc_opt=paths=source_relative \
		proto/inventory.proto
