package sql

import (
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/janithl/kottu2020/domain/blog"
	"github.com/jmoiron/sqlx"
)

type blogRepository struct {
	db database
}

func (b *blogRepository) Store(blog *blog.Blog) error {
	return nil
}

func (b *blogRepository) Find(id int) (*blog.Blog, error) {
	br := BlogRow{}
	query := sq.Select("bid, blogName, blogURL, blogRSS").
		From("blogs").Where(sq.Eq{"bid": id})

	err := b.db.Get(&br, query)
	if err != nil {
		return nil, blog.ErrNotFound
	}

	return blog.NewBlog(br.BID, br.BlogName, br.BlogURL, br.BlogRSS), nil
}

// Count returns the total number of blogs
func (b *blogRepository) Count() int {
	query := sq.Select("count(*)").
		From("blogs").
		Where(sq.Eq{"active": 1})

	var count int
	err := b.db.Get(&count, query)
	if err != nil {
		return 0
	}

	return count
}

// FindAll returns the paginated list of all blogs
func (b *blogRepository) FindAll(limit int, page int) []*blog.Blog {
	br := []*BlogRow{}
	query := sq.Select("bid, blogName, blogURL, blogRSS").
		From("blogs").
		Where(sq.Eq{"active": 1}).
		OrderBy("blogName ASC").
		Offset(uint64((page - 1) * limit)).
		Limit(uint64(limit))

	err := b.db.Select(&br, query)
	if err != nil {
		return nil
	}

	blogs := make([]*blog.Blog, 0, len(br))
	for _, b := range br {
		blog := blog.NewBlog(b.BID, b.BlogName, b.BlogURL, b.BlogRSS)
		blogs = append(blogs, blog)
	}

	return blogs
}

// FindAll returns the list of most popular blogs
func (b *blogRepository) FindPopular(limit int) []*blog.Blog {
	br := []*BlogRow{}
	query := sq.Select("bid, blogName, blogURL, blogRSS").
		From("blogs").
		Where(sq.Eq{"active": 1}).
		OrderBy("blogName ASC").
		Limit(uint64(limit))
	// TODO: order by popularity

	err := b.db.Select(&br, query)
	if err != nil {
		return nil
	}

	blogs := make([]*blog.Blog, 0, len(br))
	for _, b := range br {
		blog := blog.NewBlog(b.BID, b.BlogName, b.BlogURL, b.BlogRSS)
		blogs = append(blogs, blog)
	}

	return blogs
}

// NewBlogRepository returns a new instance of a blog repository
func NewBlogRepository(conn string) blog.Repository {
	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	return &blogRepository{
		db: database{db: db},
	}
}
