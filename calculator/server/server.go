package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/sum/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const addr = "0.0.0.0:50051"

type server struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	pb.RegisterSumServiceServer(s, &server{})

	log.Println("Server is running on ", addr)
	log.Fatalln(s.Serve(lis))
}
