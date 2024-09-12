package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func doGetBlogs(ctx context.Context, c pb.BlogServiceClient) []*pb.Blog {
	stream, err := c.GetBlogs(ctx, &emptypb.Empty{})

	if err != nil {
		log.Fatalf("%v", err)
	}

	blogs := make([]*pb.Blog, 0)

	for {
		blog, err := stream.Recv()

		if err == io.EOF {
			return blogs
		}

		if err != nil {
			log.Printf("Unexpected error in strean. Error: %v\n", err)
			continue
		}

		blogs = append(blogs, blog)
	}
}

func doCreateBlog(ctx context.Context, c pb.BlogServiceClient, blog *pb.Blog) *pb.BlogId {
	newBlog, err := c.CreateBlog(ctx, blog)

	if err != nil {
		log.Fatalf("%v", err)
	}

	return newBlog
}

func doGetBlogById(ctx context.Context, c pb.BlogServiceClient, blogId *pb.BlogId) *pb.Blog {
	blog, err := c.GetBlogById(ctx, blogId)

	if err != nil {
		log.Fatalf("%v", err)
	}

	return blog
}

func doUpdateBlog(ctx context.Context, c pb.BlogServiceClient, blog *pb.Blog) *pb.Blog {
	_, err := c.UpdateBlog(ctx, blog)

	if err != nil {
		log.Fatalf("%v", err)
	}

	return blog
}

func doDeleteBlogById(ctx context.Context, c pb.BlogServiceClient, blogId *pb.BlogId) *pb.BlogId {
	_, err := c.DeleteBlogById(ctx, blogId)

	if err != nil {
		log.Fatalf("%v", err)
	}

	return blogId
}
