package types

type UsersService interface {
	ValidateCredentials(email, password string) error
	CheckUserExistence(email string) error
	GetHashedPassword(password string) (string, error)
	GetTokenForUser(password, hashedPassword string) (string, error)
}
