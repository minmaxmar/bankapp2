package handlers

import (
	"bankapp2/app/models"
	"bankapp2/restapi/operations"

	"github.com/rs/zerolog/log"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetCards(params operations.GetCardsParams) middleware.Responder {
	log.Info().Msgf("Trying to GET cards from storage\n")

	ctx := params.HTTPRequest.Context()
	cards, err := h.controller.GetCards(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to GET cards from storage")
		return operations.NewGetCardsDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET Cards from storage " + err.Error(),
			},
		})
	}

	return operations.NewGetCardsOK().WithPayload(cards)
}

func (h *handlers) GetCardID(params operations.GetCardsIDParams) middleware.Responder {
	log.Info().Msgf("Trying to GET card from storage, card id: %+v\n", convertI64tStr(params.ID))

	if params.ID == 0 {
		return operations.NewGetCardsIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET card from storage, card id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	card, err := h.controller.GetCardID(ctx, int64(params.ID))

	if err != nil {
		log.Error().
			Err(err).
			Int("ID:", params.ID). // Use Int(), Float64(), Str(), Interface()  etc. for different types
			Msg("Failed to GET Card from storage")
		return operations.NewGetCardsIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET card from storage, card id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewGetCardsIDOK().WithPayload(&card)
}

func (h *handlers) DeleteCardsID(params operations.DeleteCardsIDParams) middleware.Responder {
	log.Info().
		Int("ID:", params.ID).
		Msg("Trying to DELETE card from storage")

	if params.ID == 0 {
		return operations.NewDeleteCardsIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE card from storage, card id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	err := h.controller.DeleteCardID(ctx, int(params.ID))

	if err != nil {
		log.Error().
			Err(err).
			Int("ID:", params.ID).
			Msg("Failed to DELETE card from storage")
		return operations.NewDeleteCardsIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE card from storage, card id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewDeleteCardsIDNoContent()
}

func (h *handlers) PostCards(params operations.PostCardsParams) middleware.Responder {
	log.Info().
		Interface("card", params.Card). // Use Interface() for complex types
		Msg("Trying to POST card in storage")

	err := validate.Struct(params.Card)
	if err != nil {
		log.Error().
			Err(err).
			Interface("card", params.Card).
			Msg("Failed to POST card in storage")
		return operations.NewGetCardsIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST card in storage " + err.Error(),
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	card, err := h.controller.PostCard(ctx, *params.Card)

	if err != nil {
		log.Error().
			Err(err).
			Interface("card", params.Card).
			Msg("Failed to POST card in storage")
		return operations.NewGetCardsIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST card in storage " + err.Error(),
			},
		})
	}

	return operations.NewPostCardsCreated().WithPayload(&card)
}
