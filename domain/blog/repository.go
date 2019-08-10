package blog

import "errors"

// Repository provides access a blog store
type Repository interface {
	Store(blog *Blog) error
	Find(id int) (*Blog, error)
	StorePost(post *Post) error
	FindPost(id int) (*Post, error)
	FindLatestPosts(language string, limit int) []*Post
}

// ErrBlogNotFound is used when a blog could not be found.
var ErrBlogNotFound = errors.New("Blog not found")

// ErrPostNotFound is used when a post could not be found.
var ErrPostNotFound = errors.New("Post not found")
