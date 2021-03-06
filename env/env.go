package env

import (
	"os"

	"github.com/joho/godotenv"
)

// when running using SUPERVISOR, os.Getenv("GOPATH") not work, change to $GOPATH manually
func init() {
	if err := godotenv.Load(os.Getenv("GOPATH") + "/src/github.com/mrizkip/backend-finding-dosen/.env"); err != nil {
		panic(err)
	}
}

// Getenv for get environment variables
func Getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
