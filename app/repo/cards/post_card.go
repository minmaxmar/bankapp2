package cards_repo

import (
	"bankapp2/app/models"
	"context"
	"log/slog"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func (repo *cardRepo) PostCard(ctx context.Context, cardData models.Card) (models.Card, error) {
	args := pgx.NamedArgs{
		"id":          cardData.ID,
		"userID":      cardData.UserID,
		"bankID":      cardData.BankID,
		"number":      cardData.Number,
		"create_date": time.Time(cardData.CreateDate),
	}
	var createTime time.Time
	card := models.Card{}
	err := repo.db.
		GetConn().
		QueryRow(ctx, postCardQuery, args).
		Scan(&card.ID, &card.UserID, &card.BankID, &card.Number, &createTime)

	card.CreateDate = strfmt.DateTime(createTime)

	if err != nil {
		return models.Card{}, err
	}

	repo.logger.Info(
		"Success POST card from storage",
		slog.Any("ID", card.ID),
	)

	return card, nil
}

func (repo *cardRepo) PostCardMy(ctx context.Context, cardData models.Card) (models.Card, error) {

	card := models.Card{}
	err := repo.db.GetConn().WithContext(ctx).Raw(postCardQuery,
		cardData.ID,
		cardData.UserID,
		cardData.BankID,
		cardData.Number,
		time.Time(cardData.CreateDate),
	).Scan(&card).Error

	if err != nil {
		return models.Card{}, err
	}
	log.Info().Msgf("Success POST card from storage. ID: %+v\n", card.ID)

	return card, nil
}
