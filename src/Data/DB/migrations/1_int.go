package migrations

import (
	db "github.com/Mehrshadshaabani/Web-API-Clean/Data/DB"
	"github.com/Mehrshadshaabani/Web-API-Clean/Data/models"
	"github.com/Mehrshadshaabani/Web-API-Clean/config"
	"github.com/Mehrshadshaabani/Web-API-Clean/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up() {
	// Apply any changes needed for the first migration
	database := db.GetDB()
	tables := []interface{}{}
	country := models.Country{}
	if !database.Migrator().HasTable(&country) {
		tables = append(tables, &country)
	}
	city := models.City{}
	if !database.Migrator().HasTable(&city) {
		tables = append(tables, &city)
	}
	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "table created", nil)

}
