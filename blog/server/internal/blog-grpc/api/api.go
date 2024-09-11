package api

import (
	"context"
	"github.com/lapeko/udemy__grpc-golang/blog/proto"
	"github.com/lapeko/udemy__grpc-golang/blog/server/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	address       = "localhost:50051"
	reflectionRun = true
)

type Api struct {
	db             *mongo.Database
	BlogRepository *storage.BlogRepository
	proto.BlogServiceServer
}

func (a *Api) Start() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	if reflectionRun {
		reflection.Register(s)
	}

	proto.RegisterBlogServiceServer(s, a)

	log.Println("Server is running on", address)
	log.Fatalln(s.Serve(lis))
}

func (a *Api) InitStorage() {
	connectOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	client, err := mongo.Connect(context.Background(), connectOptions)
	if err != nil {
		log.Fatalln(err)
	}
	a.db = client.Database("blog-grpc")
	log.Println("mongodb successfully connected")
	a.BlogRepository = storage.NewBlogRepository(a.db)
}
