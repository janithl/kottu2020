package service

import (
	"time"

	"github.com/janithl/kottu2020/domain/blog"
)

// blogService holds the implementation of the blog service
type blogService struct {
	blogs blog.Repository
}

// StoreNewBlog stores a new blog
func (s *blogService) StoreNewBlog(name string, siteURL string, feedURL string) (*blog.Blog, error) {
	b := blog.NewBlog(0, name, siteURL, feedURL)
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

// StoreNewPost stores a new post
func (s *blogService) StoreNewPost(blogID int, title string, content string, link string) (*blog.Post, error) {
	b, err := s.FindBlog(blogID)
	if err != nil {
		return nil, err
	}

	post := blog.NewPost(0, blogID, title, content, link, time.Now(), time.Now())
	b.AddPost(post)

	err = s.blogs.Store(b)
	if err != nil {
		return nil, err
	}

	err = s.blogs.StorePost(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// FindPost finds a post by its ID
func (s *blogService) FindPost(id int) (*blog.Post, error) {
	p, err := s.blogs.FindPost(id)
	return p, err
}

// NewBlogService returns a new instance of the blog service
func NewBlogService(repo blog.Repository) blog.Service {
	return &blogService{
		blogs: repo,
	}
}
