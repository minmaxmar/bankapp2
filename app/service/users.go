package service

import (
	"bankapp2/app/models"
	"context"
)

func (s service) GetUserID(ctx context.Context, id int64) (models.User, error) {
	// here no need for transaction
	return s.userRepo.GetUserID(s.userRepo.GetConn(), ctx, id)
}

func (s service) PostUser(ctx context.Context, userdData models.NewUser) (models.User, error) {

	tx := s.userRepo.BeginTransaction()
	// tx, err := s.cardRepo.BeginTransaction() tx creation error handled where originated  - TODO: pass error!!!

	user := models.User{
		FirstName: userdData.FirstName,
		LastName:  userdData.LastName,
		Email:     userdData.Email,
	}
	// TODO: here rabbitmq was used.
	userReturn, err := s.userRepo.PostUser(tx, ctx, user) // all queries are executed in the transaction

	if err != nil {
		return models.User{}, err
	}

	defer s.RollbackOrCommit(tx, err)

	return userReturn, nil
}

func (s service) DeleteUserID(ctx context.Context, id int64) error {
	tx := s.cardRepo.BeginTransaction()
	// return s.rabbitMQ.ProduceDeleteCard(ctx, id)
	_, err := s.userRepo.DeleteUserID(tx, ctx, id) // rowsAffected

	defer s.RollbackOrCommit(tx, err)

	return err
}

func (s service) GetUsers(ctx context.Context) ([]*models.User, error) {
	return s.userRepo.GetUsers(s.userRepo.GetConn(), ctx)
}
