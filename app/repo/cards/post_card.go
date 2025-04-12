package cards_repo

import (
	"bankapp2/app/models"
	repoModels "bankapp2/app/repo"
	"context"
	"log/slog"
	"time"

	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
)

func (repo *cardRepo) PostCard(connWithOrNoTx *gorm.DB, ctx context.Context, cardData models.Card) (models.Card, error) {

	card := repoModels.Card{
		CardNumber: cardData.Number,
		ExpireDate: cardData.ExpiryDate,
		Total:      cardData.Total,
		BankID:     cardData.BankID,
		UserID:     cardData.UserID,
		CreateDate: strfmt.DateTime(time.Now()),
	}

	if err := connWithOrNoTx.WithContext(ctx).Create(&card).Error; err != nil {
		return models.Card{}, err
	}

	repo.logger.Info(
		"Success POST card from storage",
		slog.Any("ID", card.ID),
	)

	return *repo.convertModel(card), nil
}
