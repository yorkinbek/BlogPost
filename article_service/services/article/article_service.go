package article

import (
	"context"
	"log"

	blogpost "yorqinbek/microservices/Blogpost/article_service/protogen/blogpost"
	"yorqinbek/microservices/Blogpost/article_service/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type articleService struct {
	stg storage.StorageI
	blogpost.UnimplementedArticleServiceServer
}

// NewArticleService ...
func NewArticleService(stg storage.StorageI) *articleService {
	return &articleService{
		stg: stg,
	}
}

// Ping ...
func (s *articleService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong{
		Message: "OK",
	}, nil
}

// CreateArticle ...
func (s *articleService) CreateArticle(ctx context.Context, req *blogpost.CreateArticleRequest) (*blogpost.Article, error) {
	id := uuid.New()

	err := s.stg.AddArticle(id.String(), req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddArticle: %s", err.Error())
	}

	article, err := s.stg.GetArticleByID(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.Stg.GetArticleByID: %s", err.Error())
	}

	return &blogpost.Article{
		Id:        article.Id,
		Content:   article.Content,
		AuthorId:  article.Author.Id,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}

// UpdateArticle ....
func (s *articleService) UpdateArticle(ctx context.Context, req *blogpost.UpdateArticleRequest) (*blogpost.Article, error) {
	err := s.stg.UpdateArticle(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateArticle: %s", err.Error())
	}

	article, err := s.stg.GetArticleByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleByID: %s", err.Error())
	}

	return &blogpost.Article{
		Id:        article.Id,
		Content:   article.Content,
		AuthorId:  article.Author.Id,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}

// DeleteArticle ....
func (s *articleService) DeleteArticle(ctx context.Context, req *blogpost.DeleteArticleRequest) (*blogpost.Article, error) {
	article, err := s.stg.GetArticleByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleByID: %s", err.Error())
	}

	err = s.stg.DeleteArticle(article.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteArticle: %s", err.Error())
	}

	return &blogpost.Article{
		Id:        article.Id,
		Content:   article.Content,
		AuthorId:  article.Author.Id,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}, nil
}

// GetArticleList ....
func (s *articleService) GetArticleList(ctx context.Context, req *blogpost.GetArticleListRequest) (*blogpost.GetArticleListResponse, error) {
	res, err := s.stg.GetArticleList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteArticle: %s", err.Error())
	}

	return res, nil
}

// GetArticleByID ....
func (s *articleService) GetArticleByID(ctx context.Context, req *blogpost.GetArticleByIDRequest) (*blogpost.GetArticleByIDResponse, error) {
	article, err := s.stg.GetArticleByID(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleByID: %s", err.Error())
	}

	return article, nil
}
