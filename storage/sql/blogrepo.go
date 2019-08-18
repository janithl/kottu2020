package sql

import (
	"log"
	"time"

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
		return nil, blog.ErrBlogNotFound
	}

	return blog.NewBlog(br.BID, br.BlogName, br.BlogURL, br.BlogRSS), nil
}

func (b *blogRepository) StorePost(post *blog.Post) error {
	return nil
}

func (b *blogRepository) FindPost(id int) (*blog.Post, error) {
	query := sq.Select("postID, blogID, title, postContent, link, serverTimestamp, api_ts, " +
		"language, fbCount, postBuzz, trend").
		From("posts").
		Where(sq.Eq{"postID": id})

	pr := PostRow{}
	err := b.db.Get(&pr, query)
	if err != nil {
		return nil, blog.ErrPostNotFound
	}

	post := blog.NewPost(pr.PostID, pr.BlogID, pr.Title, pr.PostContent, pr.Link,
		time.Unix(pr.ServerTimestamp, 0), time.Unix(pr.APITimestamp, 0))

	post.SetLanguage(pr.Language)
	post.SetStatistics(int(pr.FBCount), pr.PostBuzz, pr.Trend)
	post.CalculateChilies()

	return post, nil
}

func (b *blogRepository) FindLatestPosts(language string, limit int, page int) []*blog.Post {
	query := sq.Select("postID, blogID, title, postContent, link, serverTimestamp, api_ts, " +
		"language, fbCount, postBuzz, trend").
		From("posts").
		Where(sq.Eq{"language": language}).
		OrderBy("serverTimestamp DESC").
		Offset(uint64((page - 1) * limit)).
		Limit(uint64(limit))

	pr := []PostRow{}
	err := b.db.Select(&pr, query)
	if err != nil {
		return nil
	}

	posts := []*blog.Post{}
	for _, p := range pr {
		post := blog.NewPost(p.PostID, p.BlogID, p.Title, p.PostContent, p.Link,
			time.Unix(p.ServerTimestamp, 0), time.Unix(p.APITimestamp, 0))
		post.SetLanguage(p.Language)
		post.SetStatistics(int(p.FBCount), p.PostBuzz, p.Trend)
		post.CalculateChilies()
		posts = append(posts, post)
	}

	return posts
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
