package services

type usersService struct{}

func NewUsersService() *usersService {
	return &usersService{}
}
func (h *usersService) ValidateCredentials(email, password string) error {
	return nil
}
func (h *usersService) CheckUserExistence(email string) error {
	return nil
}
func (h *usersService) GetHashedPassword(password string) (string, error) {
	return "", nil
}
func (h *usersService) GetTokenForUser(password, hashedPassword string) (string, error) {
	return "", nil
}
