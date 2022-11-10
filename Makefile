run:
	go run main.go

gen:
	protoc --go_out=proto/ --go_opt=paths=source_relative \
    --go-grpc_out=proto/ --go-grpc_opt=paths=source_relative \
    --proto_path=proto/ proto/*.proto