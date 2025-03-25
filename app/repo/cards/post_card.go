package cards_repo

import (
	"bankapp2/app/models"
	repoModels "bankapp2/app/repo"
	"context"
	"log/slog"

	"gorm.io/gorm"
)

// func (repo *cardRepo) PostCard0(connWithOrNoTx *gorm.DB, ctx context.Context, cardData models.Card) (models.Card, error) {

// 	// card := models.Card{}
// 	card := repoModels.Card{}
// 	err := connWithOrNoTx.WithContext(ctx).Raw(postCardQuery,
// 		11,
// 		cardData.UserID,
// 		cardData.BankID,
// 		cardData.Number,
// 		time.Time(cardData.CreateDate),
// 		cardData.ExpiryDate,
// 		cardData.Total,
// 	).Scan(&card).Error
// 	// Scan(&card.ID, &card.UserID, &card.BankID, &card.Number, &card.CreateDate)

// 	if err != nil {
// 		return models.Card{}, err
// 	}
// 	repo.logger.Info(
// 		"Success POST card from storage",
// 		slog.Any("ID", card.ID),
// 	)

// 	return models.Card{
// 		BankID: card.BankID,
// 	}, nil
// }

func (repo *cardRepo) PostCard(connWithOrNoTx *gorm.DB, ctx context.Context, cardData models.Card) (models.Card, error) {

	card := repoModels.Card{
		CardNumber: cardData.Number,
		ExpireDate: cardData.ExpiryDate,
		Total:      cardData.Total,
		BankID:     cardData.BankID,
		ClientID:   cardData.UserID,
	}

	repo.logger.Info(
		"TRYYYYYYYYYING!",
		slog.Any("card", card),
	)

	if err := connWithOrNoTx.WithContext(ctx).Create(&card).Error; err != nil {
		return models.Card{}, err
	}

	repo.logger.Info(
		"Success POST card from storage",
		slog.Any("ID", card.ID),
	)

	returnModel := *repo.modelConv(card)

	return returnModel, nil
}
