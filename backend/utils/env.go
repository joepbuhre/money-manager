package utils

import (
	"log"
	"money-manager/database"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type Env struct {
	DB      *pgxpool.Pool
	Queries *database.Queries
}

func GetConfig(key string) string {

	// name of config file (without extension)
	viper.SetConfigFile(".env")
	// look for config in the working directory
	viper.AddConfigPath(".")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
