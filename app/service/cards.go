package service

import (
	"bankapp2/app/models"
	"context"
	"time"

	"github.com/go-openapi/strfmt"
)

func (s service) GetCardID(ctx context.Context, id int64) (models.Card, error) {
	// here no need for transaction
	return s.cardRepo.GetCardID(s.cardRepo.GetConn(), ctx, id)
}

func (s service) PostCard(ctx context.Context, cardData models.NewCard) (models.Card, error) {

	tx := s.cardRepo.BeginTransaction()
	// tx, err := s.cardRepo.BeginTransaction() tx creation error handled where originated

	// id, _ := uuid.NewUUID()

	//TODO: check USER exist!!! throuh userRepo
	// TODO: bank exists!

	// TODO : to fix "Failed to POST card in storage ERROR: column \"bank_id\" of relation \"cards\" does not exist (SQLSTATE 42703
	// implement all TODOs above to validate user bank existing and fill in card models with rhis information

	card := models.Card{
		// ID:         int64(id.ID()),
		UserID:     cardData.UserID,
		BankID:     cardData.BankID,
		Number:     cardData.Number,
		CreateDate: strfmt.DateTime(time.Now()),
		ExpiryDate: cardData.ExpiryDate,
		Total:      0,
	}
	// TODO: here rabbitmq was used.
	cardReturn, err := s.cardRepo.PostCard(tx, ctx, card) // all queries are executed in the transaction

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
