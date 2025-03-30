package service

import (
	"bankapp2/app/models"
	"context"
)

func (s service) GetBankID(ctx context.Context, id int64) (models.Bank, error) {
	return s.bankRepo.GetBankID(s.bankRepo.GetConn(), ctx, id)
}

func (s service) PostBank(ctx context.Context, bankData models.NewBank) (models.Bank, error) {

	tx := s.bankRepo.BeginTransaction()
	// tx, err := s.cardRepo.BeginTransaction() tx creation error handled where originated  - TODO: pass error!!!

	bank := models.Bank{
		Name: bankData.Name,
	}
	// TODO: here rabbitmq was used.
	bankReturn, err := s.bankRepo.PostBank(tx, ctx, bank) // all queries are executed in the transaction

	if err != nil {
		return models.Bank{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	return bankReturn, nil
}

func (s service) DeleteBankID(ctx context.Context, id int64) error {
	tx := s.bankRepo.BeginTransaction()
	// return s.rabbitMQ.ProduceDeleteCard(ctx, id)
	_, err := s.bankRepo.DeleteBankID(tx, ctx, id) // rowsAffected

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	return err
}

func (s service) GetBanks(ctx context.Context) ([]*models.Bank, error) {
	return s.bankRepo.GetBanks(s.bankRepo.GetConn(), ctx)
}
