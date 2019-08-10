package sql

import (
	"database/sql"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/janithl/kottu2020/domain/blog"
)

type blogRepository struct {
	db *sql.DB
}

func (b *blogRepository) Store(blog *blog.Blog) error {
	return nil
}

func (b *blogRepository) Find(id int) (*blog.Blog, error) {
	blogQuery := sq.Select("bid, blogName, blogURL, blogRSS").From("blogs").Where(squirrel.Eq{"bid": id})

	var br blogRow
	err := blogQuery.RunWith(b.db).QueryRow().Scan(&br.bid, &br.blogName, &br.blogURL, &br.blogRSS)
	if err != nil {
		return nil, blog.ErrBlogNotFound
	}

	return blog.NewBlog(br.bid, br.blogName, br.blogURL, br.blogRSS), nil
}

func (b *blogRepository) StorePost(post *blog.Post) error {
	return nil
}

func (b *blogRepository) FindPost(id int) (*blog.Post, error) {
	postQuery := sq.Select("postID, blogID, title, postContent, link, serverTimestamp, api_ts").From("posts").Where(squirrel.Eq{"postID": id})

	var pr postRow
	err := postQuery.RunWith(b.db).QueryRow().Scan(&pr.postID, &pr.blogID, &pr.title,
		&pr.postContent, &pr.link, &pr.serverTimestamp, &pr.apiTs)
	if err != nil {
		return nil, blog.ErrPostNotFound
	}

	return blog.NewPost(pr.postID, pr.blogID, pr.title, pr.postContent, pr.link,
		time.Unix(pr.serverTimestamp, 0), time.Unix(pr.apiTs, 0)), nil
}

func (b *blogRepository) FindLatestPosts(language string, limit int) []*blog.Post {
	return nil
}

// NewBlogRepository returns a new instance of a blog repository
func NewBlogRepository(conn string) blog.Repository {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	return &blogRepository{
		db: db,
	}
}
