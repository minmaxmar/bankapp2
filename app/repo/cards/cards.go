package cards_repo

import (
	"bankapp2/app/models"
	"bankapp2/helper/database"
	"context"
)

const (
	getCardIDQuery    = `select * from cards where id = @cardID`
	postCardQuery     = `insert into cards (id, userid, bankid, number, created_at) values (@id, @userID, @bankID, @number, @create_date) returning *`
	deleteCardIDQuery = `delete from cards where id = @cardID`
	getCardsQuery     = `select * from cards`
)

type cardRepo struct {
	db database.DB
}

type CardsRepo interface {
	GetCardID(ctx context.Context, id int64) (models.Card, error)
	PostCard(ctx context.Context, card models.Card) (models.Card, error)
	DeleteCardID(ctx context.Context, id int64) (int64, error)
	GetCards(ctx context.Context) ([]*models.Card, error)
}

func NewCardRepo(db database.DB) CardsRepo {
	return &cardRepo{
		db: db,
	}
}
