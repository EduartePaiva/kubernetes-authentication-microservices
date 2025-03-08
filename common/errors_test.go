package common

import (
	"errors"
	"net/http"
	"testing"
)

func TestHandleHttpError(t *testing.T) {
	// test if a not http error will be wrote with those inputs
	randomError := errors.New("some error")
	w := new(HttpResponseWriterMock)
	mc := w.
		On("WriteHeader", http.StatusOK).
		On("Write", []byte("{\"error\":\""+"something went wrong"+"\"}")).Return(0, nil)
	HandleHttpError(randomError, w, http.StatusOK)
	w.AssertExpectations(t)
	mc.Unset()
	// test with a http error, it should prioritize using the status and message of the http error
	httpError := HttpError{Code: http.StatusBadRequest, Message: "Super error"}
	w.
		On("WriteHeader", http.StatusBadRequest).
		On("Write", []byte("{\"error\":\""+"Super error"+"\"}")).Return(0, nil)
	HandleHttpError(httpError, w, http.StatusOK)
	w.AssertExpectations(t)
}
