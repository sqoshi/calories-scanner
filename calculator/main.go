package main

import (
	"calculator/api"
	"calculator/database"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var wg sync.WaitGroup

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "debug"
	}
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	logrus.SetLevel(ll)
}

func main() {
	os.Setenv("DB_NAME", "fooddatabase")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_USR", "postgres")
	os.Setenv("DB_PWD", "postgres")
	wg.Add(1)
	go api.RunAPI(&wg)
	database.Connect()
	wg.Wait()
}
