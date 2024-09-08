package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	address = "0.0.0.0:50051"
)

type greetServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalln(err)
	}

	creds, err := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.key")

	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterGreetServiceServer(s, &greetServer{})

	log.Println("Server running on ", address)
	log.Fatalln(s.Serve(lis))
}
