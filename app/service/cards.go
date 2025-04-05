package service

import (
	"bankapp2/app/models"
	"context"
	"log/slog"
	"time"

	"github.com/go-openapi/strfmt"
)

func (s service) GetCardID(ctx context.Context, id int64) (models.Card, error) {
	// here no need for transaction
	return s.cardRepo.GetCardID(s.cardRepo.GetConn(), ctx, id)
}

func (s service) PostCard(ctx context.Context, cardData models.NewCard) (models.Card, error) {

	tx := s.cardRepo.BeginTransaction()
	if _, err := s.bankRepo.GetBankID(tx, ctx, cardData.BankID); err != nil {
		s.logger.Error(
			"Error fetching bank ID:",
			slog.Any("ID", cardData.BankID),
		)
		return models.Card{}, err
	}

	if _, err := s.userRepo.GetUserID(tx, ctx, cardData.UserID); err != nil {
		s.logger.Error(
			"Error fetching user ID:",
			slog.Any("ID", cardData.UserID),
		)
		return models.Card{}, err
	}

	card := models.Card{
		UserID:     cardData.UserID,
		BankID:     cardData.BankID,
		Number:     cardData.Number,
		CreateDate: strfmt.DateTime(time.Now()),
		ExpiryDate: cardData.ExpiryDate,
		Total:      0,
	}
	cardReturn, err := s.cardRepo.PostCard(tx, ctx, card)

	if err != nil {
		return models.Card{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	return cardReturn, nil
}

func (s service) DeleteCardID(ctx context.Context, id int64) error {
	tx := s.cardRepo.BeginTransaction()
	// return s.rabbitMQ.ProduceDeleteCard(ctx, id)
	_, err := s.cardRepo.DeleteCardID(tx, ctx, id) // rowsAffected

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	return err
}

func (s service) GetCards(ctx context.Context) ([]*models.Card, error) {
	return s.cardRepo.GetCards(s.cardRepo.GetConn(), ctx)
}

// func (s *service) StartTransaction(ctx context.Context) (*sql.Tx, error) {
//     return s.cardRepo.BeginTransaction()
// }
