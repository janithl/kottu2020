package click

import "time"

// Click type holds information about post clicks
type Click struct {
	PostID    int
	IPAddress string
	CreatedAt time.Time
}
