package utils

import (
	"errors"
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/types"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
)

func HandleHttpError(err error, w http.ResponseWriter, fallbackErrorCode int) {
	authErr, ok := err.(types.AuthError)
	if ok {
		common.WriteError(w, authErr.Code, authErr)
		return
	}
	common.WriteError(w, fallbackErrorCode, errors.New("something went wrong"))
}
