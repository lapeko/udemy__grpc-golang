BINARY_CWD=.\bin\calculator
SERVER_SRC_CWD=.\server
CLIENT_SRC_CWD=.\client

protoc:
	protoc -I=proto --go_out=./proto --go-grpc_out=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*.proto

build:
	if not exist $(BINARY_CWD) mkdir $(BINARY_CWD)
	go build -v -o $(BINARY_CWD) $(SERVER_SRC_CWD)
	go build -v -o $(BINARY_CWD) $(CLIENT_SRC_CWD)

run-server:
	go run -v $(SERVER_SRC_CWD)

run-client:
	go run -v $(CLIENT_SRC_CWD)