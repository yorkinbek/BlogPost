package auth

import (
	"context"
	"log"

	"yorqinbek/microservices/Blogpost/auth_service/config"
	blogpost "yorqinbek/microservices/Blogpost/auth_service/protogen/blogpost"
	"yorqinbek/microservices/Blogpost/auth_service/storage"
	"yorqinbek/microservices/Blogpost/auth_service/util"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authService struct {
	cfg config.Config
	stg storage.StorageI
	blogpost.UnimplementedAuthServiceServer
}

// NewAuthService ...
func NewAuthService(cfg config.Config, stg storage.StorageI) *authService {
	return &authService{
		cfg: cfg,
		stg: stg,
	}
}

// Ping ...
func (s *authService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong{
		Message: "OK",
	}, nil
}

// CreateAuth ...
func (s *authService) CreateUser(ctx context.Context, req *blogpost.CreateUserRequest) (*blogpost.User, error) {
	id := uuid.New()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "util.HashPassword: %s", err.Error())
	}

	req.Password = hashedPassword

	err = s.stg.AddUser(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddUser: %s", err.Error())
	}

	user, err := s.stg.GetUserByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.Stg.GetUserByID: %s", err.Error())
	}

	return user, nil
}

// UpdateUser ....
func (s *authService) UpdateUser(ctx context.Context, req *blogpost.UpdateUserRequest) (*blogpost.User, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "util.HashPassword: %s", err.Error())
	}

	req.Password = hashedPassword

	err = s.stg.UpdateUser(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateUser: %s", err.Error())
	}

	user, err := s.stg.GetUserByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserByID: %s", err.Error())
	}

	return user, nil
}

// DeleteUser ....
func (s *authService) DeleteUser(ctx context.Context, req *blogpost.DeleteUserRequest) (*blogpost.User, error) {
	user, err := s.stg.GetUserByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserByID: %s", err.Error())
	}

	err = s.stg.DeleteUser(user.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteUser: %s", err.Error())
	}

	return user, nil
}

// GetUserList ....
func (s *authService) GetUserList(ctx context.Context, req *blogpost.GetUserListRequest) (*blogpost.GetUserListResponse, error) {
	res, err := s.stg.GetUserList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserList: %s", err.Error())
	}

	return res, nil
}

// GetUserByID ....
func (s *authService) GetUserByID(ctx context.Context, req *blogpost.GetUserByIDRequest) (*blogpost.User, error) {
	user, err := s.stg.GetUserByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserByID: %s", err.Error())
	}

	return user, nil
}
