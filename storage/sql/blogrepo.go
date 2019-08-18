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
