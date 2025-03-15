package cards_repo

import (
	"bankapp2/app/models"
	"context"

	"gorm.io/gorm"
)

func (repo *cardRepo) GetCards(connWithOrNoTx *gorm.DB, ctx context.Context) ([]*models.Card, error) {

	var cards []*models.Card
	if err := connWithOrNoTx.
		WithContext(ctx).
		Raw(getCardsQuery).Scan(&cards).Error; err != nil {
		return nil, err
	}

	repo.logger.Info("Success GET cards from storage")

	return cards, nil
}
