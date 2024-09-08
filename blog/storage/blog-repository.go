package storage

import (
	"context"
	"github.com/lapeko/udemy__grpc-golang/blog/internal/blog-grpc/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		return nil, err
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, err
	}

	return &oid, nil
}
