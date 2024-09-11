package storage

import (
	"context"
	"fmt"
	"github.com/lapeko/udemy__grpc-golang/blog/server/internal/blog-grpc/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
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

func (br *BlogRepository) GetAll() ([]*models.Blog, error) {
	ctx := context.Background()
	cur, err := br.collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error finding documents: %v", err))
	}

	defer func(cur *mongo.Cursor, ctx context.Context) {
		if err := cur.Close(ctx); err != nil {
			log.Printf("Error closing cursor: %v", err)
		}
	}(cur, ctx)

	blogs := make([]*models.Blog, 0)

	for cur.Next(ctx) {
		var blog models.Blog
		if err := cur.Decode(&blog); err != nil {
			log.Printf("Error during decoding. Error: %v\n", err)
			continue
		}
		blogs = append(blogs, &blog)
	}

	if err := cur.Err(); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Cursos error. Error: %v\n", err))
	}

	return blogs, nil
}
