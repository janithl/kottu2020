package inmemory

import "sync"
import "github.com/janithl/kottu2020/domain/blog"

type blogRepository struct {
	mutex     sync.RWMutex
	blogIndex int
	postIndex int
	blogs     map[int]*blog.Blog
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
	return nil, blog.ErrNotFound
}

// NewBlogRepository returns a new instance of a blog repository
func NewBlogRepository() blog.Repository {
	return &blogRepository{
		blogs:     make(map[int]*blog.Blog),
		blogIndex: 0,
		postIndex: 0,
	}
}
