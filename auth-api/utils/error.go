package utils

import (
	"errors"
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/types"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
)

func HandleHttpError(err error, w http.ResponseWriter) {
	authErr, ok := err.(types.AuthError)
	if ok {
		common.WriteError(w, authErr.Code, authErr)
		return
	}
	common.WriteError(w, http.StatusInternalServerError, errors.New("something went wrong"))
}
