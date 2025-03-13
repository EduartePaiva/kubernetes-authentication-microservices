package transports

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/types"
)

var (
	authApiAddress = common.EnvString("AUTH_API_ADDRESS", "http://localhost:3000")
)

type restTransportSvc struct{}

// GetHashedPassword implements types.TransportsService.
func (r *restTransportSvc) GetHashedPassword(ctx context.Context, password string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", authApiAddress+"/hashed-pw/"+password, nil)
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to create user."}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to create user."}
	}
	body := make(map[string]string)
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to create user."}
	}
	errBody, ok := body["error"]
	if ok {
		return "", common.HttpError{Code: res.StatusCode, Message: errBody}
	}
	return body["hashed"], nil
}

// GetToken implements types.TransportsService.
func (r *restTransportSvc) GetToken(ctx context.Context, password string, hashedPassword string) (string, error) {
	data := map[string]string{
		"password":       password,
		"hashedPassword": hashedPassword,
	}
	jBytes, _ := json.Marshal(&data)
	req, err := http.NewRequestWithContext(ctx, "POST", authApiAddress+"/token", bytes.NewReader(jBytes))
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to verify user."}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to verify user."}
	}
	body := make(map[string]string)
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", common.HttpError{Code: http.StatusInternalServerError, Message: "Failed to verify user."}
	}
	errBody, ok := body["error"]
	if ok {
		return "", common.HttpError{Code: res.StatusCode, Message: errBody}
	}
	return body["token"], nil
}

// GetTokenConfirmation implements types.TransportsService.
func (r *restTransportSvc) GetTokenConfirmation(ctx context.Context, token string) (bool, error) {
	panic("unimplemented")
}

type gRPCTransportSvc struct{}

// GetHashedPassword implements types.TransportsService.
func (g *gRPCTransportSvc) GetHashedPassword(ctx context.Context, password string) (string, error) {
	panic("unimplemented")
}

// GetToken implements types.TransportsService.
func (g *gRPCTransportSvc) GetToken(ctx context.Context, password string, hashedPassword string) (string, error) {
	panic("unimplemented")
}

// GetTokenConfirmation implements types.TransportsService.
func (g *gRPCTransportSvc) GetTokenConfirmation(ctx context.Context, token string) (bool, error) {
	panic("unimplemented")
}

func NewTransportService(transportProtocol string) types.TransportsService {
	switch transportProtocol {
	case "REST":
		return &restTransportSvc{}
	case "gRPC":
		return &gRPCTransportSvc{}
	}
	panic("only 'REST' or 'gRPC' are valid communication protocol")
}
