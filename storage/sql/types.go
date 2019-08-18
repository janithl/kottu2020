package sql

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type database struct {
	db *sqlx.DB
}

func (d *database) Get(v interface{}, b squirrel.SelectBuilder) error {
	query, args, err := b.ToSql()
	if err != nil {
		return err
	}
	return d.db.Get(v, query, args...)
}

func (d *database) Select(v interface{}, b squirrel.SelectBuilder) error {
	query, args, err := b.ToSql()
	if err != nil {
		return err
	}
	return d.db.Select(v, query, args...)
}

// BlogRow holds the blog details in the db
type BlogRow struct {
	BID      int    `db:"bid"`
	BlogName string `db:"blogName"`
	BlogURL  string `db:"blogURL"`
	BlogRSS  string `db:"blogRSS"`
}

// PostRow holds the post data in the db
type PostRow struct {
	PostID          int     `db:"postID"`
	BlogID          int     `db:"blogID"`
	Title           string  `db:"title"`
	PostContent     string  `db:"postContent"`
	Link            string  `db:"link"`
	ServerTimestamp int64   `db:"serverTimestamp"`
	APITimestamp    int64   `db:"api_ts"`
	Language        string  `db:"language"`
	FBCount         int64   `db:"fbCount"`
	PostBuzz        float32 `db:"postBuzz"`
	Trend           float32 `db:"trend"`
}
