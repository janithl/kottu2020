package click

// Service provides access a click service
type Service interface {
	StoreClick(postID int, ipAddress string) (*Click, error)
}
