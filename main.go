package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/janithl/kottu2020/service"
	"github.com/janithl/kottu2020/storage/sql"
	"github.com/janithl/kottu2020/web"
)

func main() {
	blogRepo := sql.NewBlogRepository(connectionString())
	blogService := service.NewBlogService(blogRepo)

	postRepo := sql.NewPostRepository(connectionString())
	postService := service.NewPostService(postRepo)

	port := flag.Int("p", 9000, "server port")
	flag.Parse()

	server := web.NewServer(*port, blogService, postService)
	server.Serve()
}

func connectionString() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/kottu", dbUser, dbPass, dbHost, dbPort)
}
