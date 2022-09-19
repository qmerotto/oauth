package database

import (
	"fmt"
	"log"
	"oauth/common/database/migrations"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(user string, password string, host string, port string, database string, poolSize int, logLevel logger.LogLevel) {
	var err error
	DB, err = gorm.Open(
		postgres.Open(fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, database)),
		&gorm.Config{
			Logger:                 logger.Default.LogMode(logLevel),
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database with user %s, and name %s. Error: %s",
			user, database, err.Error()))
	}
	setConnectionPool(DB, poolSize)
}

func RunMigrations() {
	m := gormigrate.New(DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		&migrations.M20220905185106,
	})
	if err := m.Migrate(); err != nil {
		log.Fatalf("db migration error: %v", err)
	}
}

func setConnectionPool(db *gorm.DB, poolSize int) {
	rawDb, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("error when setting connection pool: %s", err))
	}
	//rawDb.SetMaxIdleConns(poolSize / 2)
	rawDb.SetMaxOpenConns(poolSize)
}
