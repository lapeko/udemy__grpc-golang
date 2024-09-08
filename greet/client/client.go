package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

const (
	address = "0.0.0.0:50051"
	useSSL  = false
)

func main() {

	var opts []grpc.DialOption

	if useSSL {
		creds, err := credentials.NewClientTLSFromFile("ssl/server.crt", "")

		if err != nil {
			log.Fatalln(err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	con, err := grpc.NewClient(address, opts...)

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
