package storage

import "go.mongodb.org/mongo-driver/mongo"

type BlogRepository struct {
	Collection *mongo.Collection
}

func NewBlogRepository(db *mongo.Database) *BlogRepository {
	return &BlogRepository{Collection: db.Collection("blogs")}
}
