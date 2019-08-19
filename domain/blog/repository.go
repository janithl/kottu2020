package blog

import "errors"

// Repository provides access a blog store
type Repository interface {
	Store(blog *Blog) error
	Find(id int) (*Blog, error)
	Count() int
	FindAll(limit int, page int) []*Blog
	FindPopular(limit int) []*Blog
}

// ErrNotFound is used when a blog could not be found.
var ErrNotFound = errors.New("Blog not found")
