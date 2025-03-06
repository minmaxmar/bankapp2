package controller

import (
	"bankapp2/app/models"
	"context"
)

func (c controller) GetCardID(ctx context.Context, id int64) (models.Card, error) {
	return c.service.GetCardID(ctx, id)
}

func (c controller) PostCard(ctx context.Context, cardData models.NewCard) (models.Card, error) {
	return c.service.PostCard(ctx, cardData)
}

func (c controller) DeleteCardID(ctx context.Context, id int64) error {
	return c.service.DeleteCardID(ctx, id)
}

func (c controller) GetCards(ctx context.Context) ([]*models.Card, error) {
	return c.service.GetCards(ctx)
}
