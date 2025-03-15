package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	authSv := NewAuthService()

	assert.NotPanics(t, func() {
		authSv.CreateToken()
	}, "test if createToken don't panic")
}
