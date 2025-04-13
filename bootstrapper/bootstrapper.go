package bootstrapper

import (
	controller "bankapp2/app/controllers"
	"bankapp2/app/handlers"
	cards_repo "bankapp2/app/repo/cards"
	kafka "bankapp2/app/repo/kafkaa"
	"bankapp2/helper/config"
	"bankapp2/helper/database"
	logger "bankapp2/helper/logger"

	banks_repo "bankapp2/app/repo/banks"
	users_repo "bankapp2/app/repo/users"
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
	Controller     controller.Controller
	Config         *config.Config
	Handlers       handlers.Handlers
	UserRepository users_repo.UsersRepo
	CardRepository cards_repo.CardsRepo
	BankRepository banks_repo.BanksRepo
	Kafka          kafka.Kafka
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r.Infrastructure.Logger = logger.NewLogger()

	r.registerRepositoriesAndServices(ctx, r.Infrastructure.DB)
	err := r.registerAPIServer(*r.Config)
	if err != nil {
		log.Fatal("cant start server")
	}

	// <-ctx.Done()
	// log.Println("Exited cleanly.")

	return nil
}

func (r *RootBootstrapper) registerRepositoriesAndServices(ctx context.Context, db database.DB) {
	logger := r.Infrastructure.Logger
	r.Infrastructure.DB = database.NewDB().NewConn(*r.Config, logger)
	r.UserRepository = users_repo.NewUsersRepo(r.Infrastructure.DB, logger)
	r.CardRepository = cards_repo.NewCardRepo(r.Infrastructure.DB, logger)
	r.BankRepository = banks_repo.NewBanksRepo(r.Infrastructure.DB, logger)

	// consumer - cron, producer - goroutine with ticker
	kp, err := kafka.NewConn(r.CardRepository, *r.Config, logger)
	if err != nil {
		log.Fatal("cant start kafka producer")
	}
	r.Kafka = kp

	if err := r.Kafka.ScheduleProducer(ctx); err != nil {
		log.Fatal(err)
	}

	if err := r.Kafka.ScheduleConsumer(ctx); err != nil {
		log.Fatal(err)
	}
	r.Service = service.New(logger, r.UserRepository, r.CardRepository, r.BankRepository, r.Kafka)
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
	// if err := r.Validator.RegisterValidation("expiry_date_validator", validators.ValidateExpiryDate); err != nil {
	// 	log.Fatal("Failed to register custom validator: " + err.Error())
	// }

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
