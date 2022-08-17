package sql

import (
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetMustEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Printf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func GetUrl() (string, error) {
	dbUser := GetMustEnv("DB_USER")
	dbPwd := GetMustEnv("DB_PASS")
	dbName := GetMustEnv("DB_NAME")
	dbPort := GetMustEnv("DB_PORT")
	instanceHost := GetMustEnv("DB_INSTANCE_HOST")
	if dbUser == "" || dbPwd == "" || dbName == "" {
		return "", errors.New("environment variables required for database connection are missing")
	}

	dbUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPwd, instanceHost, dbPort, dbName,
	)
	return dbUrl, nil
}
