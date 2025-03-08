package common

type HttpError struct {
	// Http status code
	Code int
	// The error message
	Message string
}

func (e HttpError) Error() string {
	return e.Message
}
