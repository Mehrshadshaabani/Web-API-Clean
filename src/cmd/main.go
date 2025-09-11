package main

import (
	"github.com/Mehrshadshaabani/Web-API-Clean/Data/catch"
	"github.com/Mehrshadshaabani/Web-API-Clean/api"
	"github.com/Mehrshadshaabani/Web-API-Clean/config"
)

func main() {
	cfg := config.GetConfig()
	catch.InitRedis(cfg)
	defer catch.CloseRedisClient()
	api.InitServer()
}
