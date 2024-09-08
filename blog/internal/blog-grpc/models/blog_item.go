package models

import (
	pb "github.com/lapeko/udemy__grpc-golang/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	Id       primitive.ObjectID `bson:"id,omitempty"`
	AuthorId string             `bson:"authorId"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func (b *Blog) toProto() *pb.Blog {
	return &pb.Blog{
		Id:       b.Id.Hex(),
		AuthorId: b.AuthorId,
		Title:    b.Title,
		Content:  b.Content,
	}
}
