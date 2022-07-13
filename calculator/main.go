package main

import (
	"calculator/api"
	"calculator/database"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	os.Setenv("DB_NAME", "fooddatabase")
	os.Setenv("DB_HOST", "0.0.0.0")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_USR", "postgres")
	os.Setenv("DB_PWD", "postgres")
	wg.Add(1)
	go api.RunAPI(&wg)
	database.Connect()
	wg.Wait()
}
