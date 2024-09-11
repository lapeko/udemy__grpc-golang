package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "0.0.0.0:50051"

func main() {
	cc, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln(err)
	}

	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(cc)

	client := pb.NewBlogServiceClient(cc)

	blogs, err := doGetAll(client)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(blogs)
}
