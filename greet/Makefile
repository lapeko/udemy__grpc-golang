BINARY_CWD=.\bin\greet
SERVER_SRC_CWD=.\server
CLIENT_SRC_CWD=.\client
SSL_CWD=.\ssl

protoc:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*.proto

build:
	if not exist $(BINARY_CWD) mkdir $(BINARY_CWD)
	go build -v -o $(BINARY_CWD) $(SERVER_SRC_CWD)
	go build -v -o $(BINARY_CWD) $(CLIENT_SRC_CWD)

run-server:
	go run -v $(SERVER_SRC_CWD)

run-client:
	go run -v $(CLIENT_SRC_CWD)

generate-ssl:
	cd $(SSL_CWD) && openssl genrsa -out server.key 2048 && openssl req -new -x509 -key server.key -out server.crt -days 365 -config openssl.cnf

run-evans:
	evans --host localhost --port 50051 -r
