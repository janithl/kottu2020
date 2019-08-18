package click

// Repository provides access a click store
type Repository interface {
	Store(click *Click) error
}
