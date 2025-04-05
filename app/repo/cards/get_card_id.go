package cards_repo

import (
	"bankapp2/app/models"
	"context"
	"log/slog"

	"gorm.io/gorm"
)

func (repo *cardRepo) GetCardID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (models.Card, error) {

	card := models.Card{}
	err := connWithOrNoTx.WithContext(ctx).Raw(getCardIDQuery,
		id,
	).Scan(&card).Error

	if err != nil {
		return models.Card{}, err
	}
	repo.logger.Info(
		"Success GET card from storage",
		slog.Any("ID", card.ID),
	)

	return card, nil
}
