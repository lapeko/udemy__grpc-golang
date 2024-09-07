package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	address = "0.0.0.0:50051"
)

type greetServer struct {
	pb.GreetServiceServer
}

type greetListServer struct {
	pb.GreetListServiceServer
}

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &greetServer{})
	pb.RegisterGreetListServiceServer(s, &greetListServer{})

	log.Println("Server running on ", address)
	log.Fatalln(s.Serve(lis))
}
