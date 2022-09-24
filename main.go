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
		common.Settings().DB().User,
		common.Settings().DB().Password,
		common.Settings().DB().Host,
		common.Settings().DB().Port,
		common.Settings().DB().Name,
		common.Settings().DB().PoolSize,
		logger.Info,
	)
	database.RunMigrations()
}

func main() {
	web_api.RunServer()
}
