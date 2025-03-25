package controller

import (
	"bankapp2/app/models"
	"context"
)

func (c controller) GetUserID(ctx context.Context, id int64) (models.User, error) {
	return c.service.GetUserID(ctx, id)
}

func (c controller) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	return c.service.PostUser(ctx, userData)
}

func (c controller) DeleteUserID(ctx context.Context, id int64) error {
	return c.service.DeleteUserID(ctx, id)
}

func (c controller) GetUsers(ctx context.Context) ([]*models.User, error) {
	return c.service.GetUsers(ctx)
}
