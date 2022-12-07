package storage

import (
	"yorqinbek/microservices/Blogpost/auth_service/protogen/blogpost"
)

// StorageI ...
type StorageI interface {
	AddUser(id string, entity *blogpost.CreateUserRequest) error
	GetUserByID(id string) (*blogpost.User, error)
	GetUserList(offset, limit int, search string) (resp *blogpost.GetUserListResponse, err error)
	UpdateUser(entity *blogpost.UpdateUserRequest) error
	DeleteUser(id string) error
	GetUserByUsername(username string) (*blogpost.User, error)
}
