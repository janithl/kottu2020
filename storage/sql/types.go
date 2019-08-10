package sql

type blogRow struct {
	bid      int
	blogName string
	blogURL  string
	blogRSS  string
}

type postRow struct {
	postID          int
	blogID          int
	title           string
	postContent     string
	link            string
	serverTimestamp int64
	apiTs           int64
}
