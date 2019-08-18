package service

import (
	"time"

	"github.com/janithl/kottu2020/domain/post"
)

// postService holds the implementation of the post service
type postService struct {
	posts post.Repository
}

// StoreNewPost stores a new post
func (s *postService) StoreNewPost(blogID int, title string, content string, link string) (*post.Post, error) {
	post := post.NewPost(0, blogID, title, content, link, time.Now(), time.Now())
	err := s.posts.Store(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// FindPost finds a post by its ID
func (s *postService) FindPost(id int) (*post.Post, error) {
	p, err := s.posts.Find(id)
	return p, err
}

// FindLatestPosts returns a list of length 'limit' of posts in given 'language'
func (s *postService) FindLatestPosts(language string, limit int, page int) []*post.Post {
	return s.posts.FindLatest(language, limit, page)
}

// NewPostService returns a new instance of the post service
func NewPostService(repo post.Repository) post.Service {
	return &postService{
		posts: repo,
	}
}
