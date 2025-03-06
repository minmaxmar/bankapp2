package cards_repo

import (
	"bankapp2/app/models"
	"context"
)

func (repo *cardRepo) DeleteCardID(ctx context.Context, id int64) (int64, error) {

	card := models.Card{ID: id}
	result := repo.db.GetConn().WithContext(ctx).Delete(&card)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
