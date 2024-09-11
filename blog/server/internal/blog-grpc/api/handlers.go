package api

import (
	"context"
	"github.com/lapeko/udemy__grpc-golang/blog/proto"
	"github.com/lapeko/udemy__grpc-golang/blog/server/internal/blog-grpc/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Api) GetBlogs(_ *emptypb.Empty, in grpc.ServerStreamingServer[proto.Blog]) error {
	return nil
}

func (a *Api) CreateBlog(c context.Context, p *proto.Blog) (*proto.BlogId, error) {
	blog := models.Blog{}
	blog.FillFromProto(p)

	oid, err := a.BlogRepository.CreateOne(blog)

	if err != nil {
		return nil, status.Error(500, "Creation failure")
	}

	return &proto.BlogId{Id: oid.Hex()}, nil
}

func (a *Api) GetBlogById(context.Context, *proto.BlogId) (*proto.Blog, error) {
	return nil, nil
}

func (a *Api) UpdateBlog(context.Context, *proto.Blog) (*emptypb.Empty, error) {
	return nil, nil
}

func (a *Api) DeleteBlog(context.Context, *proto.Blog) (*proto.BlogId, error) {
	return nil, nil
}
