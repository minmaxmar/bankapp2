package database

import (
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB interface {
	NewConn(dsn string) DB
	GetConn() *gorm.DB
}

type db struct {
	conn *gorm.DB
}

func (d *db) NewConn(dsn string) DB {

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
		conn: conn,
	}
}

func (db *db) GetConn() *gorm.DB {
	return db.conn
}
