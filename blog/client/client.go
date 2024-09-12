package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "0.0.0.0:50051"

func main() {
	cc, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln(err)
	}

	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(cc)

	client := pb.NewBlogServiceClient(cc)
	c := context.Background()

	blog := pb.Blog{
		AuthorId: "Test Id",
		Content:  "Test content",
		Title:    "Test title",
	}

	newBLog := doCreateBlog(c, client, &blog)
	log.Printf("New blog created with id: %s\n", newBLog.Id)
	allBlogs := doGetBlogs(c, client)
	log.Printf("newBlog.Id equals allBlogs[last].Id: %t\n", newBLog.Id == allBlogs[len(allBlogs)-1].Id)
	allBlogs[len(allBlogs)-1].Content = "Updated content"
	doUpdateBlog(c, client, allBlogs[len(allBlogs)-1])
	res := doGetBlogById(c, client, newBLog)
	log.Printf("Content updated successfully: %t\n", res.Content == "Updated content")
	doDeleteBlogById(c, client, newBLog)
	oldSize := len(allBlogs)
	allBlogs = doGetBlogs(c, client)
	log.Printf("Blog deleted successfully: %t\n", len(allBlogs) == oldSize-1)
}
