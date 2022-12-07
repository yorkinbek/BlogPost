package storage

import (
	"yorqinbek/microservices/Blogpost/article_service/protogen/blogpost"
)

// StorageI ...
type StorageI interface {
	AddArticle(id string, entity *blogpost.CreateArticleRequest) error
	GetArticleByID(id string) (*blogpost.GetArticleByIDResponse, error)
	GetArticleList(offset, limit int, search string) (resp *blogpost.GetArticleListResponse, err error)
	UpdateArticle(entity *blogpost.UpdateArticleRequest) error
	DeleteArticle(id string) error

	GetAuthorByID(id string) (*blogpost.Author, error)
	AddAuthor(id string, entity *blogpost.CreateAuthorRequest) error
	GetAuthorList(offset, limit int, search string) (*blogpost.GetAuthorListResponse, error)
	UpdateAuthor(entity *blogpost.UpdateAuthorRequest) error
	DeleteAuthor(id string) error
}
