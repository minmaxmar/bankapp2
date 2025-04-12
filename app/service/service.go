//go:generate mockgen -source=./service.go -destination=../mocks/service_mock.go -package=mock
package service

import (
	"bankapp2/app/models"
	cards_repo "bankapp2/app/repo/cards"
	"log/slog"

	banks_repo "bankapp2/app/repo/banks"
	kafka "bankapp2/app/repo/kafkaa"
	users_repo "bankapp2/app/repo/users"
	"context"
)

type service struct {
	logger *slog.Logger

	userRepo users_repo.UsersRepo
	bankRepo banks_repo.BanksRepo
	cardRepo cards_repo.CardsRepo

	kafka kafka.Kafka
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

	GetBankID(ctx context.Context, id int64) (models.Bank, error)
	PostBank(ctx context.Context, user models.NewBank) (models.Bank, error)
	DeleteBankID(ctx context.Context, id int64) error
	GetBanks(ctx context.Context) ([]*models.Bank, error)
}

func New(
	logger *slog.Logger,
	userRepo users_repo.UsersRepo,
	cardRepo cards_repo.CardsRepo,
	bankRepo banks_repo.BanksRepo,
	kafka kafka.Kafka,
) Service {
	return &service{
		logger:   logger,
		userRepo: userRepo,
		cardRepo: cardRepo,
		bankRepo: bankRepo,
		kafka:    kafka,
	}
}
