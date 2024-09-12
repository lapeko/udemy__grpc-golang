package api

import (
	"context"
	"fmt"
	pb "github.com/lapeko/udemy__grpc-golang/blog/proto"
	"github.com/lapeko/udemy__grpc-golang/blog/server/internal/blog-grpc/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Api) CreateBlog(c context.Context, p *pb.Blog) (*pb.BlogId, error) {
	blog := models.Blog{}
	blog.FillFromProto(p)

	oid, err := a.BlogRepository.CreateOne(c, blog)

	if err != nil {
		return nil, status.Error(codes.Internal, "Creation failure")
	}

	return &pb.BlogId{Id: oid.Hex()}, nil
}

func (a *Api) GetBlogs(_ *emptypb.Empty, stream grpc.ServerStreamingServer[pb.Blog]) error {
	blogs, err := a.BlogRepository.GetAll(context.Background())

	if err != nil {
		return err
	}

	for _, blog := range blogs {
		if err := stream.Send(blog.ToProto()); err != nil {
			return status.Error(codes.Internal, fmt.Sprintf("Stream send error. Error: %v\n", err))
		}
	}

	return nil
}

func (a *Api) GetBlogById(c context.Context, blogId *pb.BlogId) (*pb.Blog, error) {
	blog, err := a.BlogRepository.GetById(c, blogId.Id)

	if err != nil {
		return nil, err
	}

	return blog.ToProto(), nil
}

func (a *Api) UpdateBlog(c context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	blog := models.Blog{}
	blog.FillFromProto(in)

	if err := a.BlogRepository.Update(c, &blog); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (a *Api) DeleteBlogById(c context.Context, in *pb.BlogId) (*pb.BlogId, error) {
	if err := a.BlogRepository.Delete(c, in.Id); err != nil {
		return nil, err
	}
	return in, nil
}
