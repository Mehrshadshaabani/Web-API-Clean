package db

import (
	"fmt"
	"log"

	"github.com/Mehrshadshaabani/Web-API-Clean/config"
	"github.com/Mehrshadshaabani/Web-API-Clean/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	var err error
	logger := logging.NewLogger(cfg)
	cnn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.Port, cfg.Postgres.SSLMode)
	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	sqlDb, _ := dbClient.DB()
	if err := sqlDb.Ping(); err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime)
	log.Println("DB connections established")
	return nil
}
func GetDB() *gorm.DB {
	return dbClient
}
func CloseDB() {
	con, _ := dbClient.DB()
	con.Close()
}
