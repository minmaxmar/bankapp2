package handlers

import (
	controller "bankapp2/app/controllers"
	"bankapp2/restapi/operations"
	"log/slog"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-playground/validator/v10"
)

type handlers struct {
	logger     *slog.Logger
	controller controller.Controller
}

var validate *validator.Validate

type Handlers interface {
	GetUsersID(params operations.GetUsersIDParams) middleware.Responder
	PostUsers(params operations.PostUsersParams) middleware.Responder
	DeleteUsersID(params operations.DeleteUsersIDParams) middleware.Responder
	GetUsers(params operations.GetUsersParams) middleware.Responder

	GetCardID(params operations.GetCardsIDParams) middleware.Responder
	PostCards(params operations.PostCardsParams) middleware.Responder
	DeleteCardsID(params operations.DeleteCardsIDParams) middleware.Responder
	GetCards(params operations.GetCardsParams) middleware.Responder

	Link(api *operations.Bankapp2API)
}

func New(controller controller.Controller, validator *validator.Validate, logger *slog.Logger) Handlers {
	validate = validator
	return &handlers{
		logger:     logger,
		controller: controller,
	}
}

func (h *handlers) Link(api *operations.Bankapp2API) {
	api.GetUsersHandler = operations.GetUsersHandlerFunc(h.GetUsers)
	api.GetUsersIDHandler = operations.GetUsersIDHandlerFunc(h.GetUsersID)
	api.PostUsersHandler = operations.PostUsersHandlerFunc(h.PostUsers)
	api.DeleteUsersIDHandler = operations.DeleteUsersIDHandlerFunc(h.DeleteUsersID)

	api.GetCardsHandler = operations.GetCardsHandlerFunc(h.GetCards)
	api.GetCardsIDHandler = operations.GetCardsIDHandlerFunc(h.GetCardID)
	api.PostCardsHandler = operations.PostCardsHandlerFunc(h.PostCards)
	api.DeleteCardsIDHandler = operations.DeleteCardsIDHandlerFunc(h.DeleteCardsID)
}

func convertI64tStr(integer int64) string {
	return strconv.FormatInt(integer, 10)
}
