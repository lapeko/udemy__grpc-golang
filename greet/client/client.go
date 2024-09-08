package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

const address = "0.0.0.0:50051"

func main() {
	creds, err := credentials.NewClientTLSFromFile("ssl/server.crt", "")

	if err != nil {
		log.Fatalln(err)
	}

	con, err := grpc.NewClient(address, grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalln(err)
	}

	defer func(con *grpc.ClientConn) {
		err := con.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(con)

	client := pb.NewGreetServiceClient(con)
	//doGreat(client)
	//doGreetList(client)
	//doGreetLong(client)
	//doGreetEveryone(client, []string{"Maria", "Anna", "Karina"})
	doDeadlineGreet(client, 4*time.Second)
}
