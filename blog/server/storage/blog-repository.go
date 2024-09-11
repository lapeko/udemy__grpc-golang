package storage

import (
	"context"
	"errors"
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

func (br *BlogRepository) CreateOne(ctx context.Context, blog models.Blog) (*primitive.ObjectID, error) {
	res, err := br.collection.InsertOne(ctx, blog)

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Can not convert ObjectID: %v", err))
	}

	return &oid, nil
}

func (br *BlogRepository) GetAll(ctx context.Context) ([]*models.Blog, error) {
	cur, err := br.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error finding documents: %v", err))
	}

	defer func() {
		if err := cur.Close(ctx); err != nil {
			log.Printf("Error closing cursor: %v", err)
		}
	}()

	blogs := make([]*models.Blog, 0)

	for cur.Next(ctx) {
		var blog models.Blog
		if err := cur.Decode(&blog); err != nil {
			log.Printf("Error during decoding. Error: %v", err)
			continue
		}
		blogs = append(blogs, &blog)
	}

	if err := cur.Err(); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Cursos error. Error: %v", err))
	}

	return blogs, nil
}

func (br *BlogRepository) GetById(ctx context.Context, id string) (*models.Blog, error) {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Provided ObjectId hex is incorrect")
	}

	res := br.collection.FindOne(ctx, bson.M{"_id": oid})

	var blog models.Blog
	err = res.Decode(&blog)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, status.Error(codes.NotFound, "Blog not found")
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("Decoding error. Error: %v", err))
	}

	return &blog, nil
}

func (br *BlogRepository) Update(ctx context.Context, blog *models.Blog) error {
	res, err := br.collection.UpdateByID(ctx, blog.Id, blog)

	if err != nil {
		return status.Error(codes.Internal, fmt.Sprintf("Update Blog DB error occurred. Error: %v"))
	}

	if res.MatchedCount == 0 {
		return status.Error(codes.NotFound, "Blog not found")
	}

	return nil
}

func (br *BlogRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return status.Error(codes.InvalidArgument, "Provided ObjectId hex is incorrect")
	}

	res, err := br.collection.DeleteOne(ctx, oid)

	if err != nil {
		return status.Error(codes.Internal, fmt.Sprintf("Delete Blog DB error occurred. Error: %v"))
	}

	if res.DeletedCount == 0 {
		return status.Error(codes.NotFound, "Blog not found")
	}

	return nil
}
