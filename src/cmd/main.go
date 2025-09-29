package main

import (
	"fmt"

	db "github.com/Mehrshadshaabani/Web-API-Clean/Data/DB"
	"github.com/Mehrshadshaabani/Web-API-Clean/Data/DB/migrations"
	"github.com/Mehrshadshaabani/Web-API-Clean/Data/catch"
	"github.com/Mehrshadshaabani/Web-API-Clean/api"
	"github.com/Mehrshadshaabani/Web-API-Clean/config"
	"github.com/Mehrshadshaabani/Web-API-Clean/pkg/logging"
)

func main() {
	cfg := *config.GetConfig()
	logger := logging.NewLogger(&cfg)
	fmt.Println("Starting main...")
	fmt.Println("Initializing Redis...")
	err := catch.InitRedis(&cfg)
	if err != nil {
		fmt.Println("Redis init failed:", err)
		panic(err)
	}
	fmt.Println("Redis initialized")
	logger.Info(logging.General, logging.Api, "Redis initialized", nil)
	fmt.Println("Initializing DB...")
	err = db.InitDb(&cfg)
	if err != nil {
		fmt.Println("DB init failed:", err)
		panic(err)
	}
	migrations.Up()
	fmt.Println("DB initialized")

	fmt.Println("Starting server...")
	api.InitServer(&cfg)
}
