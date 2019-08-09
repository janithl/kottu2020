package blog

import "time"

// Post type holds information about the posts in blogs
type Post struct {
	ID              int
	BlogID          int
	Title           string
	Content         string
	Link            string
	Language        string
	CreatedAtRemote time.Time
	CreatedAtLocal  time.Time
	ShareCount      int
	PostBuzz        float32
	Trend           float32
}

// NewPost creates a new blog post
func NewPost(id int, blogID int, title string, content string, link string,
	createdAtRemote time.Time, createdAtLocal time.Time) *Post {
	return &Post{
		ID:              id,
		BlogID:          blogID,
		Title:           title,
		Content:         content,
		Link:            link,
		Language:        "en",
		CreatedAtRemote: createdAtRemote,
		CreatedAtLocal:  createdAtLocal,
		ShareCount:      0,
		PostBuzz:        0.0,
		Trend:           0.0,
	}
}
