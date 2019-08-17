package inmemory

import "sync"
import "github.com/janithl/kottu2020/domain/blog"

type blogRepository struct {
	mutex     sync.RWMutex
	blogIndex int
	postIndex int
	blogs     map[int]*blog.Blog
	posts     map[int]*blog.Post
}

// Store implements the storage of blogs into memory
func (r *blogRepository) Store(b *blog.Blog) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if b.ID == 0 {
		r.blogIndex++
		b.ID = r.blogIndex
	}
	r.blogs[b.ID] = b
	return nil
}

// Find finds a blog from the in-memory repo using an ID
func (r *blogRepository) Find(id int) (*blog.Blog, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if val, ok := r.blogs[id]; ok {
		return val, nil
	}
	return nil, blog.ErrBlogNotFound
}

// StorePost implements the storage of posts into memory
func (r *blogRepository) StorePost(p *blog.Post) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if p.ID == 0 {
		r.postIndex++
		p.ID = r.postIndex
	}
	r.posts[p.ID] = p
	return nil
}

// FindPost finds a post from the in-memory repo using an ID
func (r *blogRepository) FindPost(id int) (*blog.Post, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if val, ok := r.posts[id]; ok {
		return val, nil
	}
	return nil, blog.ErrPostNotFound
}

// FindLatest returns the latest posts up to a limit from the in-memory repo
func (r *blogRepository) FindLatestPosts(language string, limit int, page int) []*blog.Post {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	posts := make([]*blog.Post, 0)
	for _, post := range r.posts {
		if len(posts) < limit {
			posts = append(posts, post)
		}
	}

	return posts
}

// NewBlogRepository returns a new instance of a blog repository
func NewBlogRepository() blog.Repository {
	return &blogRepository{
		blogs:     make(map[int]*blog.Blog),
		posts:     make(map[int]*blog.Post),
		blogIndex: 0,
		postIndex: 0,
	}
}
