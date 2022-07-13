package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var Connection *sql.DB

func getEnv(key string) string {
	val, present := os.LookupEnv(key)
	if !present {
		log.Fatalf("Missing variable %s, can not connect to database \n", key)
	}
	return val
}

func buildConnectionString() string {
	return strings.Join([]string{"sslmode=disable",
		fmt.Sprintf("user=%s", getEnv("DB_USR")),
		fmt.Sprintf("password=%s", getEnv("DB_PWD")),
		fmt.Sprintf("dbname=%s", getEnv("DB_NAME")),
		fmt.Sprintf("host=%s", getEnv("DB_HOST")),
		fmt.Sprintf("port=%s", getEnv("DB_PORT")),
	}, " ")

}

func Connect() {
	var err error
	log.Info("Trying to connect to database")
	Connection, err = sql.Open("postgres", buildConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Successfully connected to database.")
}
