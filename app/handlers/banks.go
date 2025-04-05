package handlers

import (
	"bankapp2/app/models"
	"bankapp2/restapi/operations"
	"log/slog"

	"github.com/rs/zerolog/log"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetBanks(params operations.GetBanksParams) middleware.Responder {
	h.logger.Info("Trying to GET banks from storage\n")

	ctx := params.HTTPRequest.Context()
	banks, err := h.controller.GetBanks(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to GET banks from storage")
		return operations.NewGetBanksDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET banks from storage " + err.Error(),
			},
		})
	}

	return operations.NewGetBanksOK().WithPayload(banks)
}

func (h *handlers) GetBanksID(params operations.GetBanksIDParams) middleware.Responder {
	h.logger.Info("Trying to GET bank from storage, bank id: " + convertI64tStr(params.ID))

	if params.ID == 0 {
		return operations.NewGetBanksIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET bank from storage, bank id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	bank, err := h.controller.GetBankID(ctx, int64(params.ID))

	if err != nil {
		h.logger.Error(
			"Failed to GET Bank from storage",
			slog.String("ID", convertI64tStr(params.ID)),
			slog.String("error", err.Error()),
		)
		return operations.NewGetBanksIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET bank from storage, user id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewGetBanksIDOK().WithPayload(&bank)
}

func (h *handlers) DeleteBanksID(params operations.DeleteBanksIDParams) middleware.Responder {
	h.logger.Info("Trying to DELETE bank from storage, bank id: " + convertI64tStr(params.ID))

	if params.ID == 0 {
		return operations.NewDeleteBanksIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE bank from storage, bank id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	err := h.controller.DeleteBankID(ctx, int64(params.ID))

	if err != nil {
		h.logger.Error(
			"Failed to DELETE bank from storage",
			slog.String("ID", convertI64tStr(params.ID)),
			slog.String("error", err.Error()),
		)
		return operations.NewDeleteBanksIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE bank from storage, bank id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewDeleteBanksIDNoContent()
}

func (h *handlers) PostBanks(params operations.PostBanksParams) middleware.Responder {
	h.logger.Info(
		"Trying to POST bank in storage",
		slog.Any("bank", params.Bank),
	)

	err := validate.Struct(params.Bank)
	if err != nil {
		h.logger.Error(
			"Failed to Validate bank",
			slog.Any("bank", params.Bank),
			slog.String("error", err.Error()),
		)
		return operations.NewGetBanksIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST bank in storage " + err.Error(),
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	bank, err := h.controller.PostBank(ctx, *params.Bank)

	if err != nil {
		h.logger.Error(
			"Failed to POST! bank in storage",
			slog.Any("user", params.Bank),
			slog.String("error", err.Error()),
		)
		return operations.NewGetBanksIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST bank in storage " + err.Error(),
			},
		})
	}

	return operations.NewPostBanksCreated().WithPayload(&bank)
}
