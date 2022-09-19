package common

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type settings struct {
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabasePoolSize int
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

func Settings() settings {
	return currentSettings
}

func (s *settings) setFromEnvVariables() error {
	var emptyVariables []string

	dbPoolSize, err := strconv.Atoi(os.Getenv("DATABASE_POOL_SIZE"))
	if err != nil {
		panic(err)
	}

	s.DatabaseName = os.Getenv("DATABASE_NAME")
	s.DatabaseUser = os.Getenv("DATABASE_USER")
	s.DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	s.DatabaseHost = os.Getenv("DATABASE_HOST")
	s.DatabasePort = os.Getenv("DATABASE_PORT")
	s.DatabasePoolSize = dbPoolSize

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
