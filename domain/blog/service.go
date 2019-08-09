package blog

// Service provides access a blog service
type Service interface {
	StoreNewBlog(name string, siteURL string, feedURL string) (*Blog, error)
	FindBlog(id int) (*Blog, error)
	StoreNewPost(blogID int, title string, content string, link string) (*Post, error)
	FindPost(id int) (*Post, error)
}
