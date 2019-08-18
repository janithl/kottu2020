package post

// Service provides access a post service
type Service interface {
	StoreNewPost(blogID int, title string, content string, link string) (*Post, error)
	FindPost(id int) (*Post, error)
	FindLatestPosts(language string, page int) []*Post
	FindPopularPosts(language string, period string, page int) []*Post
	FindBlogPosts(blogID int, sort string, page int) []*Post
	SearchPosts(query string, language string, page int) []*Post
}

// new service
// public function insertclick($ip, $pid)
