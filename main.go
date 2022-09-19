package main

import (
	"oauth/common"
	"oauth/common/database"
	"oauth/web_api"

	"gorm.io/gorm/logger"
)

func init() {
	common.InitializeConfiguration()

	database.Connect(
		common.Settings().DatabaseUser,
		common.Settings().DatabasePassword,
		common.Settings().DatabaseHost,
		common.Settings().DatabasePort,
		common.Settings().DatabaseName,
		common.Settings().DatabasePoolSize,
		logger.Info,
	)
	database.RunMigrations()
}

func main() {
	web_api.RunServer()
}
