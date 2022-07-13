package computer

import (
	db "calculator/database"
	"calculator/types"
	log "github.com/sirupsen/logrus"
)

func ComputeCalories(foodList types.FoodList) int {
	rows, err := db.Connection.Query("SELECT * FROM food")
	log.Info(err)
	log.Info(rows)
	return 1
}
