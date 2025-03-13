package handlers

import (
	"context"

	"github.com/EduartePaiva/kubernetes-authentication-microservices/auth-api/types"
	pb "github.com/EduartePaiva/kubernetes-authentication-microservices/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcHandler struct {
	pb.UnimplementedAuthServiceServer

	service types.AuthService
}

func NewGRPCHandler(grpcServer *grpc.Server, service types.AuthService) {
	handler := &grpcHandler{service: service}
	pb.RegisterAuthServiceServer(grpcServer, handler)
}

func (h *grpcHandler) GetHashedPassword(context.Context, *pb.GetHashedPasswordReq) (*pb.GetHashedPasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHashedPassword not implemented")
}
func (h *grpcHandler) GetToken(context.Context, *pb.GetTokenReq) (*pb.GetTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (h *grpcHandler) GetTokenConfirmation(context.Context, *pb.GetTokenConfirmationReq) (*pb.GetTokenConfirmationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenConfirmation not implemented")
}
