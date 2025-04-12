package cards_repo

import (
	"bankapp2/app/models"
	repoModels "bankapp2/app/repo"
	"context"

	"gorm.io/gorm"
)

func (repo *cardRepo) GetCards(connWithOrNoTx *gorm.DB, ctx context.Context) ([]*models.Card, error) {
	var cards []*repoModels.Card
	if err := connWithOrNoTx.WithContext(ctx).Find(&cards).Error; err != nil {
		return nil, err
	}
	repo.logger.Info("Success GET cards from storage")

	returnModels := make([]*models.Card, len(cards))
	for i, card := range cards {
		returnModels[i] = repo.convertModel(*card)
	}

	return returnModels, nil
}
