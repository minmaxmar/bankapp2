package cards_repo

import (
	"bankapp2/app/models"
	"context"

	"gorm.io/gorm"
)

func (repo *cardRepo) DeleteCardID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (int64, error) {

	card := models.Card{ID: id}
	result := connWithOrNoTx.WithContext(ctx).Delete(&card)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
