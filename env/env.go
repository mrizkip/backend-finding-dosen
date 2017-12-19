package env

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(os.Getenv("GOPATH") + "/src/github.com/mrizkip/backend-finding-dosen/.env"); err != nil {
		panic(err)
	}
}

func Getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
