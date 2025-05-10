package handlers

import (
	"bankapp2/app/models"
	"bankapp2/restapi/operations"
	"log/slog"

	"github.com/rs/zerolog/log"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetUsers(params operations.GetUsersParams) middleware.Responder {
	h.logger.Info("Trying to GET users from storage\n")

	ctx := params.HTTPRequest.Context()
	users, err := h.controller.GetUsers(ctx)

	if err != nil {
		log.Error().Err(err).Msg("Failed to GET users from storage")
		return operations.NewGetUsersDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET users from storage " + err.Error(),
			},
		})
	}

	return operations.NewGetUsersOK().WithPayload(users)
}

func (h *handlers) GetUsersID(params operations.GetUsersIDParams) middleware.Responder {
	h.logger.Info("Trying to GET user from storage, user id: " + convertI64tStr(params.ID))

	if params.ID == 0 {
		return operations.NewGetUsersIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET user from storage, user id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	user, err := h.controller.GetUserID(ctx, int64(params.ID))

	if err != nil {
		h.logger.Error(
			"Failed to GET User from storage",
			slog.String("ID", convertI64tStr(params.ID)),
			slog.String("error", err.Error()),
		)
		return operations.NewGetUsersIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET user from storage, user id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewGetUsersIDOK().WithPayload(&user)
}

func (h *handlers) DeleteUsersID(params operations.DeleteUsersIDParams) middleware.Responder {
	h.logger.Info("Trying to DELETE user from storage, user id: " + convertI64tStr(params.ID))

	if params.ID == 0 {
		return operations.NewDeleteUsersIDDefault(404).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE user from storage, user id = 0",
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	err := h.controller.DeleteUserID(ctx, int64(params.ID))

	if err != nil {
		h.logger.Error(
			"Failed to DELETE user from storage",
			slog.String("ID", convertI64tStr(params.ID)),
			slog.String("error", err.Error()),
		)
		return operations.NewDeleteUsersIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE user from storage, user id: " + convertI64tStr(params.ID) + " " + err.Error(),
			},
		})
	}

	return operations.NewDeleteUsersIDNoContent()
}

func (h *handlers) PostUsers(params operations.PostUsersParams) middleware.Responder {
	h.logger.Info(
		"Trying to POST user in storage",
		slog.Any("user", params.User),
	)

	err := validate.Struct(params.User)
	if err != nil {
		h.logger.Error(
			"Failed to Validate user",
			slog.Any("user", params.User),
			slog.String("error", err.Error()),
		)
		return operations.NewGetUsersIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST user in storage " + err.Error(),
			},
		})
	}

	ctx := params.HTTPRequest.Context()
	user, err := h.controller.PostUser(ctx, *params.User)

	if err != nil {
		h.logger.Error(
			"Failed to POST! user in storage",
			slog.Any("user", params.User),
			slog.String("error", err.Error()),
		)
		return operations.NewGetUsersIDDefault(500).WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to POST user in storage " + err.Error(),
			},
		})
	}

	return operations.NewPostUsersCreated().WithPayload(&user)
}
