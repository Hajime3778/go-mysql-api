package main

import (
	"go-mysql-api/cmd/server"
	"go-mysql-api/pkg/infrastructure/config"
	"go-mysql-api/pkg/infrastructure/database"
	"go-mysql-api/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	utils.LoggingSetting()
	cfg := config.NewConfig()
	db := database.NewDB(cfg)

	server := server.NewServer(cfg, db)
	server.SetUpRouter()
	server.Run()
}
