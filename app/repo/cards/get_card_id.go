package cards_repo

import (
	"bankapp2/app/models"
	repoModels "bankapp2/app/repo"
	"context"
	"log/slog"

	"gorm.io/gorm"
)

func (repo *cardRepo) GetCardID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (models.Card, error) {

	card := repoModels.Card{}
	err := connWithOrNoTx.WithContext(ctx).First(&card, id).Error

	if err != nil {
		return models.Card{}, err
	}
	repo.logger.Info(
		"Success GET card from storage",
		slog.Any("ID", card.ID),
	)

	return *repo.convertModel(card), nil
}
