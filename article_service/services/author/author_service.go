package author

import (
	"context"
	"log"

	blogpost "yorqinbek/microservices/Blogpost/article_service/protogen/blogpost"
	"yorqinbek/microservices/Blogpost/article_service/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authorService struct {
	stg storage.StorageI
	blogpost.UnimplementedAuthorServiceServer
}

// NewAuthorService ...
func NewAuthorService(stg storage.StorageI) *authorService {
	return &authorService{
		stg: stg,
	}
}

// Ping ...
func (s *authorService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong{
		Message: "OK",
	}, nil
}

// CreateAuthor ...
func (s *authorService) CreateAuthor(ctx context.Context, req *blogpost.CreateAuthorRequest) (*blogpost.Author, error) {
	id := uuid.New()

	err := s.stg.AddAuthor(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddAuthor: %s", err.Error())
	}

	author, err := s.stg.GetAuthorByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.Stg.GetAuthorByID: %s", err.Error())
	}

	return &blogpost.Author{
		Id:        author.Id,
		Fullname:  author.Fullname,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}, nil
}

// UpdateAuthor ....
func (s *authorService) UpdateAuthor(ctx context.Context, req *blogpost.UpdateAuthorRequest) (*blogpost.Author, error) {
	err := s.stg.UpdateAuthor(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateAuthor: %s", err.Error())
	}

	author, err := s.stg.GetAuthorByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetAuthorByID: %s", err.Error())
	}

	return &blogpost.Author{
		Id:        author.Id,
		Fullname:  author.Fullname,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}, nil
}

// DeleteAuthor ....
func (s *authorService) DeleteAuthor(ctx context.Context, req *blogpost.DeleteAuthorRequest) (*blogpost.Author, error) {
	author, err := s.stg.GetAuthorByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetAuthorByID: %s", err.Error())
	}

	err = s.stg.DeleteAuthor(author.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteAuthor: %s", err.Error())
	}

	return &blogpost.Author{
		Id:        author.Id,
		Fullname:  author.Fullname,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}, nil
}

// GetAuthorList ....
func (s *authorService) GetAuthorList(ctx context.Context, req *blogpost.GetAuthorListRequest) (*blogpost.GetAuthorListResponse, error) {
	res, err := s.stg.GetAuthorList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteAuthor: %s", err.Error())
	}

	return res, nil
}

// GetAuthorByID ....
func (s *authorService) GetAuthorByID(ctx context.Context, req *blogpost.GetAuthorByIDRequest) (*blogpost.GetAuthorByIDResponse, error) {
	author, err := s.stg.GetAuthorByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetAuthorByID: %s", err.Error())
	}

	return author, nil
}
