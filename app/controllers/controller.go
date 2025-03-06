//go:generate mockgen -source=./controller.go -destination=../mocks/controller_mock.go -package=mock
package controller

import (
	"bankapp2/app/models"
	"bankapp2/app/service"
	"context"
)

type controller struct {
	// logger *slog.Logger
	service service.Service
}

type Controller interface {
	// GetUserID(ctx context.Context, id int) (models.User, error)
	// PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	// DeleteUserID(ctx context.Context, id int) error
	// GetUsers(ctx context.Context) ([]*models.User, error)

	GetCardID(ctx context.Context, id int64) (models.Card, error)
	PostCard(ctx context.Context, user models.NewCard) (models.Card, error)
	DeleteCardID(ctx context.Context, id int64) error
	GetCards(ctx context.Context) ([]*models.Card, error)
}

func New(service service.Service) Controller {
	return &controller{
		service: service,
	}
}
