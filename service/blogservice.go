package service

import "github.com/janithl/kottu2020/domain/blog"

// blogService holds the implementation of the blog service
type blogService struct {
	blogs blog.Repository
}

// StoreNewBlog stores a new blog
func (s *blogService) StoreNewBlog(name string, siteURL string, feedURL string) (*blog.Blog, error) {
	b := blog.New(0, name, siteURL, feedURL)
	err := s.blogs.Store(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// FindBlog finds a blog by its ID
func (s *blogService) FindBlog(id int) (*blog.Blog, error) {
	b, err := s.blogs.Find(id)
	return b, err
}

// NewBlogService returns a new instance of the blog service
func NewBlogService(repo blog.Repository) blog.Service {
	return &blogService{
		blogs: repo,
	}
}
