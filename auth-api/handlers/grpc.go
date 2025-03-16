package handlers

import (
	"context"
	"log"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/types"
	"github.com/EduartePaiva/kubernetes-authentication-microservices/common"
	pb "github.com/EduartePaiva/kubernetes-authentication-microservices/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedAuthServiceServer

	service types.AuthService
}

func NewGRPCHandler(grpcServer *grpc.Server, service types.AuthService) {
	handler := &grpcHandler{service: service}
	pb.RegisterAuthServiceServer(grpcServer, handler)
}

func (h *grpcHandler) GetHashedPassword(ctx context.Context, req *pb.GetHashedPasswordReq) (*pb.GetHashedPasswordRes, error) {
	hashedPw, err := h.service.CreatePasswordHash(req.Password)
	if err != nil {
		return nil, common.ConvertHttpErrorToGrpcError(err)
	}

	return &pb.GetHashedPasswordRes{HashedPassword: hashedPw}, nil
}
func (h *grpcHandler) GetToken(ctx context.Context, req *pb.GetTokenReq) (*pb.GetTokenRes, error) {
	log.Println("getting jwt token for user")
	err := h.service.VerifyPasswordHash(req.Password, req.HashedPassword)
	if err != nil {
		return nil, common.ConvertHttpErrorToGrpcError(err)
	}
	token := h.service.CreateToken()
	return &pb.GetTokenRes{Token: token}, nil
}
func (h *grpcHandler) GetTokenConfirmation(ctx context.Context, req *pb.GetTokenConfirmationReq) (*pb.GetTokenConfirmationRes, error) {
	err := h.service.VerifyToken(req.Token)
	if err != nil {
		return nil, common.ConvertHttpErrorToGrpcError(err)
	}
	return &pb.GetTokenConfirmationRes{IsValid: true}, nil
}
