package models

import (
	pb "github.com/lapeko/udemy__grpc-golang/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Blog struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string             `bson:"authorId"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func (b *Blog) ToProto() *pb.Blog {
	return &pb.Blog{
		Id:       b.Id.Hex(),
		AuthorId: b.AuthorId,
		Title:    b.Title,
		Content:  b.Content,
	}
}

func (b *Blog) FillFromProto(p *pb.Blog) {
	id, err := primitive.ObjectIDFromHex(p.Id)
	if err != nil {
		log.Println("Object ID format error. Using default nil value")
	}
	b.Id = id
	b.AuthorId = p.AuthorId
	b.Title = p.Title
	b.Content = p.Content
}
