SERVER_CWD=.\server

.PHONY: protoc run-evans server

protoc:
	protoc -I=proto --go_out=./proto --go-grpc_out=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*.proto

run-evans:
	evans --host localhost --port 50051 -r

server:
	cd $(SERVER_CWD) && go run -v .\cmd\blog-grpc
