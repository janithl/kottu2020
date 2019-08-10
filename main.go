package main

import (
	"flag"

	"github.com/janithl/kottu2020/service"
	"github.com/janithl/kottu2020/storage/inmemory"
	"github.com/janithl/kottu2020/web"
)

func main() {
	blogRepo := inmemory.NewBlogRepository()
	blogService := service.NewBlogService(blogRepo)

	port := flag.String("p", "9000", "server port")
	flag.Parse()

	server := web.NewServer(*port, blogService)
	server.Serve()
}
