package configs

import (
	"fmt"
	"os"
)

func GetDSN() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", user, password, host, dbName)
}
