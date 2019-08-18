package inmemory

import "sync"
import "github.com/janithl/kottu2020/domain/post"

type postRepository struct {
	mutex     sync.RWMutex
	blogIndex int
	postIndex int
	posts     map[int]*post.Post
}

// Store implements the storage of posts into memory
func (r *postRepository) Store(p *post.Post) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if p.ID == 0 {
		r.postIndex++
		p.ID = r.postIndex
	}
	r.posts[p.ID] = p
	return nil
}

// Find finds a post from the in-memory repo using an ID
func (r *postRepository) Find(id int) (*post.Post, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if val, ok := r.posts[id]; ok {
		return val, nil
	}
	return nil, post.ErrNotFound
}

// FindLatest returns the latest posts up to a limit from the in-memory repo
func (r *postRepository) FindLatest(language string, limit int, page int) []*post.Post {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	posts := make([]*post.Post, 0)
	for _, post := range r.posts {
		if len(posts) < limit {
			posts = append(posts, post)
		}
	}

	return posts
}

// NewPostRepository returns a new instance of a blog repository
func NewPostRepository() post.Repository {
	return &postRepository{
		posts:     make(map[int]*post.Post),
		blogIndex: 0,
		postIndex: 0,
	}
}
