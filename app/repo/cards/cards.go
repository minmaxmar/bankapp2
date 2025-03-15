package cards_repo

import (
	"bankapp2/app/models"
	"bankapp2/helper/database"
	"context"
	"log/slog"

	"gorm.io/gorm"
)

const (
	getCardIDQuery    = `select * from cards where id = @cardID`
	postCardQuery     = `insert into cards (id, userid, bankid, number, created_at) values (@id, @userID, @bankID, @number, @create_date) returning *`
	deleteCardIDQuery = `delete from cards where id = @cardID`
	getCardsQuery     = `select * from cards`
)

type cardRepo struct {
	db     database.DB
	logger *slog.Logger
}

type CardsRepo interface {
	GetCardID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (models.Card, error)
	PostCard(connWithOrNoTx *gorm.DB, ctx context.Context, card models.Card) (models.Card, error)
	DeleteCardID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (int64, error)
	GetCards(connWithOrNoTx *gorm.DB, ctx context.Context) ([]*models.Card, error)
	BeginTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB)
	RollbackTransaction(tx *gorm.DB)
	GetConn() *gorm.DB
}

func NewCardRepo(db database.DB, logger *slog.Logger) CardsRepo {
	return &cardRepo{
		db:     db,
		logger: logger,
	}
}

func (repo *cardRepo) GetConn() *gorm.DB {
	return repo.db.GetConn()
}

func (repo *cardRepo) BeginTransaction() *gorm.DB {
	return repo.db.BeginTx()
}

func (repo *cardRepo) CommitTransaction(tx *gorm.DB) {
	repo.db.CommitTx(tx)
}

func (repo *cardRepo) RollbackTransaction(tx *gorm.DB) {
	repo.db.RollbackTx(tx)
}
