package types

type AuthService interface {
	CreatePasswordHash(password string) (string, error)
	VerifyPasswordHash(password, hashedPassword string) error
	CreateToken() string
	VerifyToken(token string) error
	// GetHashedPassword(w http.ResponseWriter, r *http.Request) // this is a middleware, it can't have transport layer things
}
