package cards_repo

import (
	"bankapp2/app/models"
	"context"

	"github.com/rs/zerolog/log"
)

func (repo *cardRepo) GetCards(ctx context.Context) ([]*models.Card, error) {

	var cards []*models.Card
	if err := repo.db.GetConn().
		WithContext(ctx).
		Raw(getCardsQuery).Scan(&cards).Error; err != nil {
		return nil, err
	}

	log.Info().Msgf("Success GET cards from storage\n")

	return cards, nil
}
