package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	address       = "0.0.0.0:50051"
	useSSL        = false
	runReflection = true
)

type greetServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalln(err)
	}

	var opts []grpc.ServerOption

	if useSSL {
		creds, err := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.key")
		if err != nil {
			log.Fatalln(err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	if runReflection {
		reflection.Register(s)
	}

	pb.RegisterGreetServiceServer(s, &greetServer{})

	log.Println("Server running on ", address)
	log.Fatalln(s.Serve(lis))
}
