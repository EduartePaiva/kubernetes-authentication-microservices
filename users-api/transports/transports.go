package transports

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	pb "github.com/EduartePaiva/kubernetes-authentication-microservices/common/api"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/users-api/types"
	"google.golang.org/grpc"
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

type gRPCTransportSvc struct {
	conn *grpc.ClientConn
}

// GetHashedPassword implements types.TransportsService.
func (g *gRPCTransportSvc) GetHashedPassword(ctx context.Context, password string) (string, error) {
	c := pb.NewAuthServiceClient(g.conn)
	res, err := c.GetHashedPassword(ctx, &pb.GetHashedPasswordReq{
		Password: password,
	})
	return res.HashedPassword, common.ConvertGrpcErrorToHttpError(err)
}

// GetToken implements types.TransportsService.
func (g *gRPCTransportSvc) GetToken(ctx context.Context, password string, hashedPassword string) (string, error) {
	c := pb.NewAuthServiceClient(g.conn)
	res, err := c.GetToken(ctx, &pb.GetTokenReq{
		Password:       password,
		HashedPassword: hashedPassword,
	})
	return res.Token, common.ConvertGrpcErrorToHttpError(err)
}

// GetTokenConfirmation implements types.TransportsService.
func (g *gRPCTransportSvc) GetTokenConfirmation(ctx context.Context, token string) (bool, error) {
	c := pb.NewAuthServiceClient(g.conn)
	res, err := c.GetTokenConfirmation(ctx, &pb.GetTokenConfirmationReq{
		Token: token,
	})
	return res.IsValid, common.ConvertGrpcErrorToHttpError(err)
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
