package services

import (
	"log"
	"net/http"
	"time"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	jtwPrivateKey = common.EnvString("JTW_PRIVATE_KEY", "insecure_key")
)

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) CreatePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Println(err)
		return "", common.HttpError{
			Code:    http.StatusInternalServerError,
			Message: "failed to create secure password.",
		}
	}
	return string(hashedPassword), nil
}
func (s *authService) VerifyPasswordHash(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	switch err {
	case nil:
		return nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return common.HttpError{
			Code:    http.StatusUnauthorized,
			Message: "Failed to verify password.",
		}
	default:
		log.Println(err)
		return common.HttpError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to verify password.",
		}
	}
}
func (s *authService) CreateToken() string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(time.Hour * 1),
		},
	})
	token, err := t.SignedString([]byte(jtwPrivateKey))
	if err != nil {
		log.Println(err)
		panic("something is wrong with your code, fix it")
	}
	return token
}
func (s *authService) VerifyToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(jtwPrivateKey), nil
	}, jwt.WithExpirationRequired(), jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		log.Println(err)
		return common.HttpError{Code: http.StatusUnauthorized, Message: "Could not verify token."}
	}
	return nil
}
