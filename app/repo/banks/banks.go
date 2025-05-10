package banks_repo

import (
	"bankapp2/app/models"
	repoModels "bankapp2/app/repo"
	"bankapp2/helper/database"
	"context"
	"log/slog"

	"gorm.io/gorm"
)

type banksRepo struct {
	db     database.DB
	logger *slog.Logger
}

type BanksRepo interface {
	GetBankID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (models.Bank, error)
	PostBank(connWithOrNoTx *gorm.DB, ctx context.Context, card models.Bank) (models.Bank, error)
	DeleteBankID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (int64, error)
	GetBanks(connWithOrNoTx *gorm.DB, ctx context.Context) ([]*models.Bank, error)
	BeginTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB)
	RollbackTransaction(tx *gorm.DB)
	GetConn() *gorm.DB
}

func NewBanksRepo(db database.DB, logger *slog.Logger) BanksRepo {
	return &banksRepo{
		db:     db,
		logger: logger,
	}
}

func (repo *banksRepo) GetConn() *gorm.DB {
	return repo.db.GetConn()
}

func (repo *banksRepo) BeginTransaction() *gorm.DB {
	return repo.db.BeginTx()
}

func (repo *banksRepo) CommitTransaction(tx *gorm.DB) {
	repo.db.CommitTx(tx)
}

func (repo *banksRepo) RollbackTransaction(tx *gorm.DB) {
	repo.db.RollbackTx(tx)
}

func (repo *banksRepo) convertModel(gormModel repoModels.Bank) *models.Bank {
	return &models.Bank{
		ID:   gormModel.ID,
		Name: gormModel.Name,
	}
}

func (repo *banksRepo) DeleteBankID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (int64, error) {

	bank := repoModels.Bank{ID: id}
	result := connWithOrNoTx.WithContext(ctx).Delete(&bank)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (repo *banksRepo) GetBankID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (models.Bank, error) {

	bank := repoModels.Bank{}

	err := connWithOrNoTx.WithContext(ctx).First(&bank, id).Error

	if err != nil {
		return models.Bank{}, err
	}

	repo.logger.Info(
		"Success GET bank from storage",
		slog.Any("ID", bank.ID),
	)

	return *repo.convertModel(bank), nil
}

func (repo *banksRepo) GetBanks(connWithOrNoTx *gorm.DB, ctx context.Context) ([]*models.Bank, error) {

	var banks []*repoModels.Bank
	if err := connWithOrNoTx.WithContext(ctx).Find(&banks).Error; err != nil {
		return nil, err
	}

	repo.logger.Info("Success GET banks from storage")

	returnModels := make([]*models.Bank, len(banks))
	for i, bank := range banks {
		returnModels[i] = repo.convertModel(*bank)
	}

	return returnModels, nil
}

func (repo *banksRepo) PostBank(connWithOrNoTx *gorm.DB, ctx context.Context, bankData models.Bank) (models.Bank, error) {

	bank := repoModels.Bank{
		Name: bankData.Name,
	}

	repo.logger.Info(
		"TRYYYYYYYYYING!",
		slog.Any("bank", bank),
	)

	if err := connWithOrNoTx.WithContext(ctx).Create(&bank).Error; err != nil {
		return models.Bank{}, err
	}

	repo.logger.Info(
		"Success POST bank from storage",
		slog.Any("ID", bank.ID),
	)

	return *repo.convertModel(bank), nil
}
