package common

import (
	"errors"
	"net/http"
)

func HandleHttpError(err error, w http.ResponseWriter, fallbackErrorCode int) {
	authErr, ok := err.(HttpError)
	if ok {
		WriteError(w, authErr.Code, authErr)
		return
	}
	WriteError(w, fallbackErrorCode, errors.New("something went wrong"))
}
