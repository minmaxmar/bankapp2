package controller

import (
	"bankapp2/app/models"
	"context"
)

func (c controller) GetBankID(ctx context.Context, id int64) (models.Bank, error) {
	return c.service.GetBankID(ctx, id)
}

func (c controller) PostBank(ctx context.Context, bankData models.NewBank) (models.Bank, error) {
	return c.service.PostBank(ctx, bankData)
}

func (c controller) DeleteBankID(ctx context.Context, id int64) error {
	return c.service.DeleteBankID(ctx, id)
}

func (c controller) GetBanks(ctx context.Context) ([]*models.Bank, error) {
	return c.service.GetBanks(ctx)
}
