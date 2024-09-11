.PHONY: protoc run-evans server client

proto:
	protoc -I=proto --go_out=./proto --go-grpc_out=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*.proto

evans:
	evans --host localhost --port 50051 -r

server:
	go run -v .\server\cmd\blog-grpc

client:
	go run -v .\client

docker:
	docker-compose up -d

docker-down:
	docker-compose down
