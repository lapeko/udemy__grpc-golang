BINARY_CWD=.\bin\sum
SERVER_SRC_CWD=.\server
CLIENT_SRC_CWD=.\client

protoc:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*.proto

build:
	if not exist $(BINARY_CWD) mkdir $(BINARY_CWD)
	go build -v -o $(BINARY_CWD) $(SERVER_SRC_CWD)
	go build -v -o $(BINARY_CWD) $(CLIENT_SRC_CWD)