//go:generate mockgen -source=./service.go -destination=../mocks/service_mock.go -package=mock
package service

import (
	"bankapp2/app/models"
	cards_repo "bankapp2/app/repo/cards"
	"log/slog"

	users_repo "bankapp2/app/repo/users"
	"context"
)

type service struct {
	logger *slog.Logger

	userRepo users_repo.UsersRepo
	cardRepo cards_repo.CardsRepo
}

type Service interface {
	GetUserID(ctx context.Context, id int64) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id int64) error
	GetUsers(ctx context.Context) ([]*models.User, error)

	GetCardID(ctx context.Context, id int64) (models.Card, error)
	PostCard(ctx context.Context, user models.NewCard) (models.Card, error)
	DeleteCardID(ctx context.Context, id int64) error
	GetCards(ctx context.Context) ([]*models.Card, error)

	// StartTransaction(ctx context.Context) (*sql.Tx, error)
}

func New(
	logger *slog.Logger,
	userRepo users_repo.UsersRepo,
	cardRepo cards_repo.CardsRepo) Service {
	return &service{

		logger: logger,

		userRepo: userRepo,
		cardRepo: cardRepo,
	}
}
