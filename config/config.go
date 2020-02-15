package config

import (
	"github.com/joho/godotenv"
	"os"
)
var Username string
var Password string
var LokiUrl string
var LokiWSUrl string
func InitEnvironmentVariables() {

	 godotenv.Load()

	Username = os.Getenv("USER_NAME")
	Password = os.Getenv("PASSWORD")
	LokiUrl = os.Getenv("LOKI_URL")
	LokiWSUrl = os.Getenv("LOKIWS_URL")
}
