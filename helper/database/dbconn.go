package database

import (
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"log/slog"
)

type DB interface {
	NewConn(dsn string, logger *slog.Logger) DB
	GetConn() *gorm.DB
	BeginTx() *gorm.DB
	CommitTx(tx *gorm.DB)
	RollbackTx(tx *gorm.DB)
}

type db struct {
	logger *slog.Logger
	conn   *gorm.DB
}

func NewDB() DB {
	return &db{}
}

func (d *db) NewConn(dsn string, slogger *slog.Logger) DB {

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		os.Exit(2)
	}

	log.Info().Msg("connected")
	conn.Logger = logger.Default.LogMode(logger.Info)

	//TODO : migrations here?

	return &db{
		logger: slogger,
		conn:   conn,
	}
}

func (db *db) GetConn() *gorm.DB {
	return db.conn
}

func (db *db) BeginTx() *gorm.DB {
	tx := db.conn.Begin()
	if tx.Error != nil {
		log.Fatal().Err(tx.Error).Msg("Failed to connect to database")
	}
	return tx
}

func (db *db) CommitTx(tx *gorm.DB) {
	if err := tx.Commit().Error; err != nil {
		log.Fatal().Err(err).Msg("Failed to commit transaction")
	}
}

func (db *db) RollbackTx(tx *gorm.DB) {
	if err := tx.Rollback().Error; err != nil {
		log.Fatal().Err(err).Msg("Failed to rollback transaction")
	}
}
