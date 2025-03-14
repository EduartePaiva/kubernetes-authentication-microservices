package common

import (
	"errors"
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	GRPC_STATUS_CODE_TO_HTTP = map[codes.Code]int{
		codes.OK:                 http.StatusOK,                  // OK
		codes.Canceled:           499,                            // CANCELLED (Client Closed Request)
		codes.Unknown:            http.StatusInternalServerError, // UNKNOWN
		codes.InvalidArgument:    http.StatusBadRequest,          // INVALID_ARGUMENT
		codes.DeadlineExceeded:   http.StatusGatewayTimeout,      // DEADLINE_EXCEEDED
		codes.NotFound:           http.StatusNotFound,            // NOT_FOUND
		codes.AlreadyExists:      http.StatusConflict,            // ALREADY_EXISTS
		codes.PermissionDenied:   http.StatusForbidden,           // PERMISSION_DENIED
		codes.ResourceExhausted:  http.StatusTooManyRequests,     // RESOURCE_EXHAUSTED
		codes.FailedPrecondition: http.StatusBadRequest,          // FAILED_PRECONDITION
		codes.Aborted:            http.StatusConflict,            // ABORTED
		codes.OutOfRange:         http.StatusBadRequest,          // OUT_OF_RANGE
		codes.Unimplemented:      http.StatusNotImplemented,      // UNIMPLEMENTED
		codes.Internal:           http.StatusInternalServerError, // INTERNAL
		codes.Unavailable:        http.StatusServiceUnavailable,  // UNAVAILABLE
		codes.DataLoss:           http.StatusInternalServerError, // DATA_LOSS
		codes.Unauthenticated:    http.StatusUnauthorized,        // UNAUTHENTICATED
	}

	HTTP_STATUS_CODE_TO_GRPC = map[int]codes.Code{
		http.StatusOK:                  codes.OK,                // 200 → Success
		499:                            codes.Canceled,          // 499 → Client Closed Request (used by some proxies)
		http.StatusInternalServerError: codes.Internal,          // 500 → Generic internal error
		http.StatusBadRequest:          codes.InvalidArgument,   // 400 → Malformed request, best fit for most cases
		http.StatusGatewayTimeout:      codes.DeadlineExceeded,  // 504 → Request took too long
		http.StatusNotFound:            codes.NotFound,          // 404 → Resource not found
		http.StatusConflict:            codes.AlreadyExists,     // 409 → Conflict when resource already exists
		http.StatusForbidden:           codes.PermissionDenied,  // 403 → No permission
		http.StatusTooManyRequests:     codes.ResourceExhausted, // 429 → Too many requests (rate limiting)
		http.StatusNotImplemented:      codes.Unimplemented,     // 501 → API not implemented
		http.StatusServiceUnavailable:  codes.Unavailable,       // 503 → Server temporarily down
		http.StatusUnauthorized:        codes.Unauthenticated,   // 401 → Authentication required
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
	httpCode, ok := GRPC_STATUS_CODE_TO_HTTP[grpcErr.Code()]
	if !ok {
		httpCode = http.StatusInternalServerError
	}
	return HttpError{
		Message: grpcErr.Message(),
		Code:    httpCode,
	}
}

func ConvertHttpErrorToGrpcError(err error) error {
	if err == nil {
		return nil
	}
	httpErr, ok := err.(HttpError)
	if !ok {
		return err
	}
	code, ok := HTTP_STATUS_CODE_TO_GRPC[httpErr.Code]
	if !ok {
		log.Println("it's better to not use a code that don't have grpc equivalent")
		return err
	}
	return status.Error(code, httpErr.Message)
}
