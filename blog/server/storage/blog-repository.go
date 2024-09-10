package storage

import (
	"context"
	"fmt"
	"github.com/lapeko/udemy__grpc-golang/blog/server/internal/blog-grpc/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BlogRepository struct {
	collection *mongo.Collection
}

func NewBlogRepository(db *mongo.Database) *BlogRepository {
	return &BlogRepository{collection: db.Collection("blogs")}
}

func (br *BlogRepository) CreateOne(blog models.Blog) (*primitive.ObjectID, error) {
	res, err := br.collection.InsertOne(context.Background(), blog)

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Can not convert ObjectID: %v", err))
	}

	return &oid, nil
}
