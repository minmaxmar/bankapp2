package bootstrapper

import (
	controller "bankapp2/app/controllers"
	"bankapp2/app/handlers"
	cards_repo "bankapp2/app/repo/cards"
	"bankapp2/helper/config"
	"bankapp2/helper/database"
	logger "bankapp2/helper/logger"
	"bankapp2/helper/validators"

	// users_repo "bankapp2/repo/users"
	"bankapp2/app/service"
	"bankapp2/restapi"
	"bankapp2/restapi/operations"
	"context"
	"log"
	"log/slog"

	"github.com/go-openapi/loads"
	"github.com/go-playground/validator/v10"
)

type RootBootstrapper struct {
	Infrastructure struct {
		Logger *slog.Logger
		Server *restapi.Server
		DB     database.DB
	}
	Controller controller.Controller
	Config     *config.Config
	Handlers   handlers.Handlers
	// UserRepository users_repo.UsersRepo
	CardRepository cards_repo.CardsRepo
	Service        service.Service

	Validator *validator.Validate
}

type RootBoot interface {
	registerRepositoriesAndServices(ctx context.Context, db database.DB)
	registerAPIServer(cfg config.Config) error
	RunAPI() error
}

func New() RootBoot {
	return &RootBootstrapper{
		Config: config.LoadConfig(),
	}
}

func (r *RootBootstrapper) RunAPI() error {
	ctx := context.Background()
	r.Infrastructure.Logger = logger.NewLogger()

	r.registerRepositoriesAndServices(ctx, r.Infrastructure.DB)
	err := r.registerAPIServer(*r.Config)
	if err != nil {
		log.Fatal("cant start server")
	}

	return nil
}

func (r *RootBootstrapper) registerRepositoriesAndServices(ctx context.Context, db database.DB) {
	logger := r.Infrastructure.Logger
	r.Infrastructure.DB = database.NewDB().NewConn(*&r.Config.DatabaseURL, logger)
	// r.UserRepository = users_repo.NewUserRepo(r.Infrastructure.DB, logger)
	r.CardRepository = cards_repo.NewCardRepo(r.Infrastructure.DB, logger)
	// r.RabbitMQ = rabbitmq.NewConn(r.UserRepository, r.CardRepository, *r.Config, logger)
	// go r.RabbitMQ.NewConsumer(ctx)
	r.Service = service.New(logger, r.CardRepository)
}

func (r *RootBootstrapper) registerAPIServer(cfg config.Config) error {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	api := operations.NewBankapp2API(swaggerSpec)

	logger := r.Infrastructure.Logger

	r.Controller = controller.New(r.Service, logger)

	r.Validator = validator.New(validator.WithRequiredStructEnabled())

	// register custom validators
	if err := r.Validator.RegisterValidation("expiry_date_validator", validators.ValidateExpiryDate); err != nil {
		log.Fatal("Failed to register custom validator: " + err.Error())
	}

	r.Handlers = handlers.New(r.Controller, r.Validator, logger)
	r.Handlers.Link(api)
	if r.Handlers == nil {
		log.Fatal("handlers initialization failed")
	}

	r.Infrastructure.Server = restapi.NewServer(api)
	r.Infrastructure.Server.Port = cfg.ServerPort
	r.Infrastructure.Server.ConfigureAPI()
	if err := r.Infrastructure.Server.Serve(); err != nil {
		log.Fatalln(err)
	}

	return nil
}
