package types

type AuthService interface {
	CreatePasswordHash(password string) error
	VerifyPasswordHash(password, hashedPassword string) error
	CreateToken() string
	VerifyToken(token string) error
	// GetHashedPassword(w http.ResponseWriter, r *http.Request) // this is a middleware, it can't have transport layer things
}

type AuthError struct {
	// Http status code
	Code int
	// The error message
	Message string
}

func (e AuthError) Error() string {
	return e.Message
}
