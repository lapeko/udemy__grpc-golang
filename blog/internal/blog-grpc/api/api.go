package api

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/blog/proto"
	"github.com/lapeko/udemy__grpc-golang/blog/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

const address = "localhost:50051"

type Api struct {
	db             *mongo.Database
	blogServer     *BlogServer
	BlogRepository *storage.BlogRepository
}

func (a *Api) Start() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	a.blogServer = &BlogServer{}

	pb.RegisterBlogServiceServer(s, a.blogServer)

	log.Println("Server is running on", address)
	log.Fatalln(s.Serve(lis))
}

func (a *Api) InitStorage() {
	connectOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username: "root",
		Password: "root",
	})
	client, err := mongo.Connect(context.Background(), connectOptions)
	if err != nil {
		log.Fatalln(err)
	}
	a.db = client.Database("blog-grpc")
	a.BlogRepository = storage.NewBlogRepository(a.db)
}
