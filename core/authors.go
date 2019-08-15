package core

import (
	"context"

	"github.com/JPZ13/service/model"
)

// AuthorsService holds methods related to authors
type AuthorsService interface {
	CreateAuthor(ctx context.Context, author *model.CreateAuthorRequest) (*model.CreateAuthorResponse, error)
}

type authorsService struct {
	baseService
}

// CreateAuthor creates an author
func (s *authorsService) CreateAuthor(ctx context.Context, author *model.CreateAuthorRequest) (*model.CreateAuthorResponse, error) {
	return nil, nil
}
