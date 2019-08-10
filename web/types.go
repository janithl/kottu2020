package web

// ServerError struct is used to convey API errors
type ServerError struct {
	ErrorCode    int
	ErrorMessage string
}
