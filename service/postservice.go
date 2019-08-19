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

// FindLatestPosts returns a paginated list of posts in given 'language'
func (s *postService) FindLatestPosts(language string, page int) []*post.Post {
	return s.posts.FindLatest(language, perPage, page)
}

// FindPopularPosts returns a paginated list of popular posts for a given language and period
func (s *postService) FindPopularPosts(language string, period string, page int) []*post.Post {
	return nil
}

// FindBlogPosts returns a paginated and sorted list of posts belonging to a given blog
func (s *postService) FindBlogPosts(blogID int, sort string, page int) []*post.Post {
	return nil
}

// SearchPosts returns a paginated list of posts for a given language matching the search query
func (s *postService) SearchPosts(query string, language string, page int) []*post.Post {
	return nil
}

// NewPostService returns a new instance of the post service
func NewPostService(repo post.Repository) post.Service {
	return &postService{
		posts: repo,
	}
}
