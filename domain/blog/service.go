package blog

import "github.com/janithl/kottu2020/domain/post"

// Service provides access a blog service
type Service interface {
	StoreNewBlog(name string, siteURL string, feedURL string) (*Blog, error)
	FindBlog(id int) (*Blog, error)
	StoreNewPost(blogID int, post *post.Post) (*post.Post, error)
}

// public function fetchnumblogs()
// public function fetchallblogs($size, $page)
// public function fetchpopblogs($size)
