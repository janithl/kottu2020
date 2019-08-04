package blog

import "time"

// Post type holds information about the posts in blogs
type Post struct {
	ID              int
	BlogID          int
	Link            string
	Title           string
	Content         string
	FeedURL         string
	Language        string
	CreatedAtRemote time.Time
	CreatedAtLocal  time.Time
	ShareCount      int
	PostBuzz        float32
	Trend           float32
}
