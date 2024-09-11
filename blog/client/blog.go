package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func doGetAll(c pb.BlogServiceClient) ([]*pb.Blog, error) {
	stream, err := c.GetBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		return nil, err
	}

	blogs := make([]*pb.Blog, 0)

	for {
		blog, err := stream.Recv()

		if err == io.EOF {
			return blogs, nil
		}

		if err != nil {
			log.Printf("Unexpected error in strean. Error: %v\n", err)
			continue
		}

		blogs = append(blogs, blog)
	}
}
