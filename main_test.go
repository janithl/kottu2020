package main_test

import (
	"testing"

	"github.com/janithl/kottu2020/domain/blog"
	"github.com/janithl/kottu2020/domain/post"
	"github.com/janithl/kottu2020/service"
	"github.com/janithl/kottu2020/storage/inmemory"
)

// TestBlogServiceBlogs adds a bunch of blogs to the repository and retrieves them
// It then adds a bunch of posts and retrieves them too
func TestBlogService(t *testing.T) {
	blogRepo := inmemory.NewBlogRepository()
	blogService := service.NewBlogService(blogRepo)

	postRepo := inmemory.NewPostRepository()
	postService := service.NewPostService(postRepo)

	// Create the blogs and see if we can retrieve them
	b1, _ := blogService.StoreNewBlog("Hello Blog", "https://blog.hello.com", "https://blog.hello.com/feed")
	b2, _ := blogService.StoreNewBlog("Mello Blog", "https://blog.mello.com", "https://blog.mello.com/feed")
	b3, _ := blogService.StoreNewBlog("Yello Blog", "https://blog.yello.com", "https://blog.yello.com/feed")

	t.Run(b1.Name, testIfFoundBlog(blogService, b1))
	t.Run(b2.Name, testIfFoundBlog(blogService, b2))
	t.Run(b3.Name, testIfFoundBlog(blogService, b3))

	// Check if blog count is 3
	count := blogService.BlogCount()
	expected := 3
	if count != expected {
		t.Errorf("Expected to find blog count %d but got %d instead!", expected, count)
	}

	// Try to find nonexistent blog, should throw correct error
	_, err := blogService.FindBlog(10)
	if err != blog.ErrNotFound {
		t.Errorf("Correct error not thrown trying to find missing blog! %s", err.Error())
	}

	// Create the posts and see if we can retrieve them
	p1, _ := postService.StoreNewPost(b1.ID, "Hello Blog Intro", "Intro", "https://blog.hello.com/post/intro")
	p2, _ := postService.StoreNewPost(b1.ID, "Hello Blog Conclusion", "Conclusion", "https://blog.hello.com/post/end")
	p3, _ := postService.StoreNewPost(b2.ID, "Mello Blog Intro", "Intro", "https://blog.mello.com/post/intro")
	p4, _ := postService.StoreNewPost(b2.ID, "Mello Blog Conclusion", "Conclusion", "https://blog.mello.com/post/end")

	t.Run(p1.Title, testIfFoundPost(postService, p1))
	t.Run(p2.Title, testIfFoundPost(postService, p2))
	t.Run(p3.Title, testIfFoundPost(postService, p3))
	t.Run(p4.Title, testIfFoundPost(postService, p4))

	// Try to find nonexistent post, should throw correct error
	_, err = postService.FindPost(10)
	if err != post.ErrNotFound {
		t.Errorf("Correct error not thrown trying to find missing post! %s", err.Error())
	}

	// Test if getting latest posts works
	posts := postService.FindLatestPosts("en", 1)
	t.Run(p1.Title, testIfPostSliceContains(posts, p1))
	t.Run(p2.Title, testIfPostSliceContains(posts, p2))
	t.Run(p3.Title, testIfPostSliceContains(posts, p3))
	t.Run(p4.Title, testIfPostSliceContains(posts, p4))
}

func testIfFoundBlog(service blog.Service, b *blog.Blog) func(*testing.T) {
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

func testIfFoundPost(service post.Service, p *post.Post) func(*testing.T) {
	return func(t *testing.T) {
		actual, err := service.FindPost(p.ID)
		if err != nil {
			t.Errorf("Error finding blog! %s", err.Error())
		}
		if actual != p {
			t.Errorf("Expected to find value %s but got %s instead!", p.Title, actual.Title)
		}
	}
}

func testIfPostSliceContains(posts []*post.Post, p *post.Post) func(*testing.T) {
	return func(t *testing.T) {
		found := false
		for _, post := range posts {
			if post == p {
				found = true
			}
		}

		if !found {
			t.Errorf("Post %s not found in slice! %v", p.Title, posts)
		}
	}
}
