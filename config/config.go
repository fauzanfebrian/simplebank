package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const DBDriver string = "postgres"

var (
	DBSource      string
	ServerAddress string
)

type loadEnv struct {
	variable     *string
	key          string
	defaultValue string
}

var mapLoadEnv = []loadEnv{
	{
		variable:     &DBSource,
		key:          "DB_SOURCE",
		defaultValue: "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable",
	},
	{
		variable:     &ServerAddress,
		key:          "SERVER_ADDRESS",
		defaultValue: ":8080",
	},
}

// Before using config variables, you should run this function first to load it.
func LoadConfig(filenames ...string) {
	err := godotenv.Load(filenames...)

	if err != nil {
		fmt.Println("err load .env file:", err)
	}

	for _, v := range mapLoadEnv {
		*v.variable = os.Getenv(v.key)
		if *v.variable == "" {
			*v.variable = v.defaultValue
		}
	}
}
