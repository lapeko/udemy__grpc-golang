package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const addr = "0.0.0.0:50051"

type sumServer struct {
	pb.SumServiceServer
}

type primesServer struct {
	pb.PrimesStreamingServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	pb.RegisterSumServiceServer(s, &sumServer{})
	pb.RegisterPrimesStreamingServiceServer(s, &primesServer{})

	log.Println("Server is running on ", addr)
	log.Fatalln(s.Serve(lis))
}
