package cards_repo

import (
	"bankapp2/app/models"
	"context"
	"log/slog"
	"time"

	"gorm.io/gorm"
)

func (repo *cardRepo) PostCard(connWithOrNoTx *gorm.DB, ctx context.Context, cardData models.Card) (models.Card, error) {

	card := models.Card{}
	err := connWithOrNoTx.WithContext(ctx).Raw(postCardQuery,
		cardData.ID,
		cardData.UserID,
		cardData.BankID,
		cardData.Number,
		time.Time(cardData.CreateDate),
	).Scan(&card).Error

	if err != nil {
		return models.Card{}, err
	}
	repo.logger.Info(
		"Success POST card from storage",
		slog.Any("ID", card.ID),
	)

	return card, nil
}
