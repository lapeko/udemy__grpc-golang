package main

import (
	"github.com/lapeko/udemy__grpc-golang/blog/server/internal/blog-grpc/api"
)

func main() {
	appApi := api.Api{}
	appApi.InitStorage()
	appApi.Start()
}
