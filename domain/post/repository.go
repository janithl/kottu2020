package post

import "errors"

// Repository provides access a post store
type Repository interface {
	Store(post *Post) error
	Find(id int) (*Post, error)
	FindLatest(language string, limit int, page int) []*Post
}

// ErrNotFound is used when a post could not be found.
var ErrNotFound = errors.New("Post not found")
