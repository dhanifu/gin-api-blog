package main

import (
	"gin-api-blog/api"
	"gin-api-blog/config"
	"gin-api-blog/data/db"
	"gin-api-blog/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := db.InitDB(cfg)
	defer db.CloseDB()

	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	api.InitServer(cfg)
}
