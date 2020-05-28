package config

import "os"

var (
	DBHost     = os.Getenv("DB_HOST")
	DBName     = os.Getenv("DB_NAME")
	DBPort     = os.Getenv("DB_PORT")
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
)
