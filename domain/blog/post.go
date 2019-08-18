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
	Chilies         int
}

// SetLanguage sets the post's language
func (p *Post) SetLanguage(language string) {
	p.Language = language
}

// SetStatistics sets the post's statistics
func (p *Post) SetStatistics(shareCount int, postBuzz float32, trend float32) {
	p.ShareCount = shareCount
	p.PostBuzz = postBuzz
	p.Trend = trend
	p.calculateChilies()
}

// calculateChilies calculates the amount of 'chilies' per post
func (p *Post) calculateChilies() {
	buzz := p.PostBuzz * 100
	if buzz <= 1 {
		p.Chilies = 1
	} else if buzz <= 15 {
		p.Chilies = 2
	} else if buzz <= 35 {
		p.Chilies = 3
	} else if buzz <= 55 {
		p.Chilies = 4
	} else {
		p.Chilies = 5
	}
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
