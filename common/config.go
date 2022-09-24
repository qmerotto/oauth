package common

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"oauth/common/database"

	"github.com/joho/godotenv"
)

type settings struct {
	Database database.Settings
}

var currentSettings settings

func InitializeConfiguration() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := currentSettings.setFromEnvVariables(); err != nil {
		panic(err)
	}
}

func (s *settings) setFromEnvVariables() error {
	var emptyVariables []string

	dbPoolSize, err := strconv.Atoi(os.Getenv("DATABASE_POOL_SIZE"))
	if err != nil {
		panic(err)
	}

	s.Database.Name = os.Getenv("DATABASE_NAME")
	s.Database.User = os.Getenv("DATABASE_USER")
	s.Database.Password = os.Getenv("DATABASE_PASSWORD")
	s.Database.Host = os.Getenv("DATABASE_HOST")
	s.Database.Port = os.Getenv("DATABASE_PORT")
	s.Database.PoolSize = dbPoolSize

	v := reflect.ValueOf(s).Elem()
	typeOfV := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			emptyVariables = append(emptyVariables, typeOfV.Field(i).Name)
		}
	}

	if len(emptyVariables) == 0 {
		return nil
	}

	return fmt.Errorf("the following environment variables are missing: %s", strings.Join(emptyVariables, ", "))
}

func Settings() settings {
	return currentSettings
}

func (s settings) DB() database.Settings {
	return s.Database
}
