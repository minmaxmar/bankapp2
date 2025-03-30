package users_repo

import (
	"bankapp2/app/models"
	repoModels "bankapp2/app/repo"
	"bankapp2/helper/database"
	"context"
	"log/slog"

	"gorm.io/gorm"
)

type usersRepo struct {
	db     database.DB
	logger *slog.Logger
}

type UsersRepo interface {
	GetUserID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (models.User, error)
	PostUser(connWithOrNoTx *gorm.DB, ctx context.Context, card models.User) (models.User, error)
	DeleteUserID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (int64, error)
	GetUsers(connWithOrNoTx *gorm.DB, ctx context.Context) ([]*models.User, error)
	BeginTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB)
	RollbackTransaction(tx *gorm.DB)
	GetConn() *gorm.DB
}

func NewUsersRepo(db database.DB, logger *slog.Logger) UsersRepo {
	return &usersRepo{
		db:     db,
		logger: logger,
	}
}

func (repo *usersRepo) GetConn() *gorm.DB {
	return repo.db.GetConn()
}

func (repo *usersRepo) BeginTransaction() *gorm.DB {
	return repo.db.BeginTx()
}

func (repo *usersRepo) CommitTransaction(tx *gorm.DB) {
	repo.db.CommitTx(tx)
}

func (repo *usersRepo) RollbackTransaction(tx *gorm.DB) {
	repo.db.RollbackTx(tx)
}

func (repo *usersRepo) modelConv(gormModel repoModels.User) (result *models.User) {
	result = &models.User{
		ID:        gormModel.ID,
		FirstName: gormModel.FirstName,
		LastName:  gormModel.LastName,
		Email:     gormModel.Email,
	}
	return
}

func (repo *usersRepo) DeleteUserID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (int64, error) {

	user := repoModels.User{ID: id}
	result := connWithOrNoTx.WithContext(ctx).Delete(&user)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (repo *usersRepo) GetUserID(connWithOrNoTx *gorm.DB, ctx context.Context, id int64) (models.User, error) {

	user := repoModels.User{}

	err := connWithOrNoTx.WithContext(ctx).First(&user, id).Error

	if err != nil {
		return models.User{}, err
	}

	repo.logger.Info(
		"Success GET user from storage",
		slog.Any("ID", user.ID),
	)

	returnModel := *repo.modelConv(user)

	return returnModel, nil
}

func (repo *usersRepo) GetUsers(connWithOrNoTx *gorm.DB, ctx context.Context) ([]*models.User, error) {

	var users []*repoModels.User
	if err := connWithOrNoTx.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	repo.logger.Info("Success GET users from storage")

	returnModels := make([]*models.User, len(users))
	for i, user := range users {
		returnModels[i] = repo.modelConv(*user)
	}

	return returnModels, nil
}

func (repo *usersRepo) PostUser(connWithOrNoTx *gorm.DB, ctx context.Context, userData models.User) (models.User, error) {

	user := repoModels.User{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
	}

	repo.logger.Info(
		"TRYYYYYYYYYING!",
		slog.Any("user", user),
	)

	if err := connWithOrNoTx.WithContext(ctx).Create(&user).Error; err != nil {
		return models.User{}, err
	}

	repo.logger.Info(
		"Success POST user from storage",
		slog.Any("ID", user.ID),
	)

	returnModel := *repo.modelConv(user)

	return returnModel, nil
}
