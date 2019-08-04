package main_test

import (
	"testing"

	"github.com/janithl/kottu2020/domain/blog"
	"github.com/janithl/kottu2020/service"
	"github.com/janithl/kottu2020/storage/inmemory"
)

func TestBlogService(t *testing.T) {
	blogRepo := inmemory.NewBlogRepository()
	blogService := service.NewBlogService(blogRepo)

	b1, _ := blogService.StoreNewBlog("Hello Blog", "https://blog.hello.com", "https://blog.hello.com/feed")
	b2, _ := blogService.StoreNewBlog("Mello Blog", "https://blog.mello.com", "https://blog.mello.com/feed")
	b3, _ := blogService.StoreNewBlog("Yello Blog", "https://blog.yello.com", "https://blog.yello.com/feed")

	t.Run(b1.Name, testIfFound(blogService, b1))
	t.Run(b2.Name, testIfFound(blogService, b2))
	t.Run(b3.Name, testIfFound(blogService, b3))
}

func testIfFound(service blog.Service, b *blog.Blog) func(*testing.T) {
	return func(t *testing.T) {
		actual, err := service.FindBlog(b.ID)
		if err != nil {
			t.Errorf("Error finding blog! %s", err.Error())
		}
		if actual != b {
			t.Errorf("Expected to find value %s but got %s instead!", b.Name, actual.Name)
		}
	}
}
