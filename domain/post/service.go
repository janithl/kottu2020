package post

// Service provides access a post service
type Service interface {
	StoreNewPost(blogID int, title string, content string, link string) (*Post, error)
	FindPost(id int) (*Post, error)
	FindLatestPosts(language string, limit int, page int) []*Post
}

// public function fetchpopularposts($lang='all', $time='off', $size, $pageno=0)
// public function fetchblogposts($blogid, $size, $pageno = 0, $sort)
// public function searchposts($str, $pageno=0, $lang='all')
// public function fetchposturl($id)

// new service
// public function insertclick($ip, $pid)
