package env

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func GanachePort() string {
	return os.Getenv("GANACHE_PORT")
}

func GanacheHost() string {
	return os.Getenv("GANACHE_HOST")
}
