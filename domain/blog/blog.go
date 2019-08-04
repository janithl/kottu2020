package blog

import "time"

// Blog type holds information about blogs
type Blog struct {
	ID       int
	Name     string
	SiteURL  string
	FeedURL  string
	PolledAt time.Time
	Active   bool
	Posts    []Post
}

// New creates a new blog with no posts
func New(id int, name string, siteURL string, feedURL string) *Blog {
	return &Blog{
		ID:       id,
		Name:     name,
		SiteURL:  siteURL,
		FeedURL:  feedURL,
		PolledAt: time.Time{},
		Active:   true,
		Posts:    make([]Post, 1),
	}
}
