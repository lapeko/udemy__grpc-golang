package api

import (
	"context"
	"github.com/lapeko/udemy__grpc-golang/blog/server/internal/blog-grpc/models"
	pb "github.com/lapeko/udemy__grpc-golang/blog/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Api) GetBlogs(_ *emptypb.Empty, in grpc.ServerStreamingServer[pb.Blog]) error {
	return nil
}

func (a *Api) CreateBlog(c context.Context, p *pb.Blog) (*pb.BlogId, error) {
	blog := models.Blog{}
	blog.FillFromProto(p)

	oid, err := a.BlogRepository.CreateOne(blog)

	if err != nil {
		return nil, status.Error(500, "Creation failure")
	}

	return &pb.BlogId{Id: oid.Hex()}, nil
}

func (a *Api) GetBlogById(context.Context, *pb.BlogId) (*pb.Blog, error) {
	return nil, nil
}

func (a *Api) UpdateBlog(context.Context, *pb.Blog) (*emptypb.Empty, error) {
	return nil, nil
}

func (a *Api) DeleteBlog(context.Context, *pb.Blog) (*pb.BlogId, error) {
	return nil, nil
}
