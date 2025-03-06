package cards_repo

import (
	"bankapp2/app/models"
	"context"

	"github.com/rs/zerolog/log"
)

func (repo *cardRepo) GetCardID(ctx context.Context, id int64) (models.Card, error) {

	card := models.Card{}
	err := repo.db.GetConn().WithContext(ctx).Raw(getCardIDQuery,
		id,
	).Scan(&card).Error

	if err != nil {
		return models.Card{}, err
	}
	log.Info().Msgf("Success GET card from storage. ID: %+v\n", card.ID)

	return card, nil
}
