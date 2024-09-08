package main

import "github.com/lapeko/udemy__grpc-golang/blog/internal/blog-grpc/api"

func main() {
	appApi := api.Api{}
	appApi.InitStorage()
	appApi.Start()
}
