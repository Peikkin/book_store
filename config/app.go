package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	dsn := "host=localhost user=root password=1234 dbname=book_store port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("не удалось подключиться к базе данных: ")
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
