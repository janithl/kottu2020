package sql

import (
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/janithl/kottu2020/domain/post"
	"github.com/jmoiron/sqlx"
)

type postRepository struct {
	db database
}

func (p *postRepository) Store(post *post.Post) error {
	return nil
}

func (p *postRepository) Find(id int) (*post.Post, error) {
	query := sq.Select("postID, blogID, title, postContent, link, serverTimestamp, api_ts, " +
		"language, fbCount, postBuzz, trend").
		From("posts").
		Where(sq.Eq{"postID": id})

	pr := PostRow{}
	err := p.db.Get(&pr, query)
	if err != nil {
		return nil, post.ErrNotFound
	}

	post := post.NewPost(pr.PostID, pr.BlogID, pr.Title, pr.PostContent, pr.Link,
		time.Unix(pr.ServerTimestamp, 0), time.Unix(pr.APITimestamp, 0))

	post.SetLanguage(pr.Language)
	post.SetStatistics(int(pr.FBCount), pr.PostBuzz, pr.Trend)
	return post, nil
}

func (p *postRepository) FindLatest(language string, limit int, page int) []*post.Post {
	query := sq.Select("postID, blogID, title, postContent, link, serverTimestamp, api_ts, " +
		"language, fbCount, postBuzz, trend").
		From("posts").
		Where(sq.Eq{"language": language}).
		OrderBy("serverTimestamp DESC").
		Offset(uint64((page - 1) * limit)).
		Limit(uint64(limit))

	pr := []PostRow{}
	err := p.db.Select(&pr, query)
	if err != nil {
		return nil
	}

	posts := []*post.Post{}
	for _, p := range pr {
		post := post.NewPost(p.PostID, p.BlogID, p.Title, p.PostContent, p.Link,
			time.Unix(p.ServerTimestamp, 0), time.Unix(p.APITimestamp, 0))
		post.SetLanguage(p.Language)
		post.SetStatistics(int(p.FBCount), p.PostBuzz, p.Trend)
		posts = append(posts, post)
	}

	return posts
}

// NewPostRepository returns a new instance of a post repository
func NewPostRepository(conn string) post.Repository {
	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	return &postRepository{
		db: database{db: db},
	}
}
