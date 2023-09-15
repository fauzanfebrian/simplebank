package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const DBDriver string = "postgres"

var (
	DBSource string
)

// Before using config variables, you should run this function first to load it.
func LoadConfig(envFullPath string) error {
	if err := godotenv.Load(envFullPath); err != nil {
		return fmt.Errorf("'%s' is not env file, please check again", envFullPath)
	}

	DBSource = os.Getenv("DB_SOURCE")

	return nil
}
