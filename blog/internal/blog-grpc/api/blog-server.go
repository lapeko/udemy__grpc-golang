package api

import pb "github.com/lapeko/udemy__grpc-golang/blog/proto"

type BlogServer struct {
	pb.BlogServiceServer
}
