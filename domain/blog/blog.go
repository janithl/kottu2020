package blog

import "time"
import "github.com/janithl/kottu2020/domain/post"

// Blog type holds information about blogs
type Blog struct {
	ID       int
	Name     string
	SiteURL  string
	FeedURL  string
	PolledAt time.Time
	Active   bool
	Posts    []post.Post
}

// AddPost adds a new post to a blog
func (b *Blog) AddPost(post *post.Post) {
	b.Posts = append(b.Posts, *post)
}

// NewBlog creates a new blog with no posts
func NewBlog(id int, name string, siteURL string, feedURL string) *Blog {
	return &Blog{
		ID:       id,
		Name:     name,
		SiteURL:  siteURL,
		FeedURL:  feedURL,
		PolledAt: time.Time{},
		Active:   true,
		Posts:    make([]post.Post, 0),
	}
}
