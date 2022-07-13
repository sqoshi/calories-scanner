package computer

import (
	db "calculator/database"
	"calculator/types"
	log "github.com/sirupsen/logrus"
)

func ComputeCalories(foodList types.FoodList) int {
	//rows, err := db.Connection.Query("SELECT table_name,table_schema FROM information_schema.tables;")
	rows, err := db.Connection.Query("SELECT name,category FROM foodschema.food;")
	log.Info(err)
	for rows.Next() {
		var table string
		var schema string
		err = rows.Scan(&table, &schema)
		if err != nil {
			log.Warning(err)
		}
		log.Println(table, schema)
	}
	return 1
}
