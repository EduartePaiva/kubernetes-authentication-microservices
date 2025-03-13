package common

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/status"
)

var (
	GRPC_STATUS_CODE_TO_HTTP = map[int]int{
		0:  http.StatusOK,                  // OK
		1:  499,                            // CANCELLED (Client Closed Request)
		2:  http.StatusInternalServerError, // UNKNOWN
		3:  http.StatusBadRequest,          // INVALID_ARGUMENT
		4:  http.StatusGatewayTimeout,      // DEADLINE_EXCEEDED
		5:  http.StatusNotFound,            // NOT_FOUND
		6:  http.StatusConflict,            // ALREADY_EXISTS
		7:  http.StatusForbidden,           // PERMISSION_DENIED
		8:  http.StatusTooManyRequests,     // RESOURCE_EXHAUSTED
		9:  http.StatusBadRequest,          // FAILED_PRECONDITION
		10: http.StatusConflict,            // ABORTED
		11: http.StatusBadRequest,          // OUT_OF_RANGE
		12: http.StatusNotImplemented,      // UNIMPLEMENTED
		13: http.StatusInternalServerError, // INTERNAL
		14: http.StatusServiceUnavailable,  // UNAVAILABLE
		15: http.StatusInternalServerError, // DATA_LOSS
		16: http.StatusUnauthorized,        // UNAUTHENTICATED
	}
)

func HandleHttpError(err error, w http.ResponseWriter, fallbackErrorCode int) {
	authErr, ok := err.(HttpError)
	if ok {
		WriteError(w, authErr.Code, authErr)
		return
	}
	WriteError(w, fallbackErrorCode, errors.New("something went wrong"))
}

func ConvertGrpcErrorToHttpError(err error) error {
	if err == nil {
		return nil
	}
	grpcErr, ok := status.FromError(err)
	if !ok {
		return HttpError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	httpCode, ok := GRPC_STATUS_CODE_TO_HTTP[int(grpcErr.Code())]
	if !ok {
		httpCode = http.StatusInternalServerError
	}
	return HttpError{
		Message: grpcErr.Message(),
		Code:    httpCode,
	}
}
