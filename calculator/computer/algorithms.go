package computer

import (
	db "calculator/database"
	"calculator/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func getDistinctNames(foodList types.FoodList) []string {
	visited := map[string]bool{}
	var result []string
	for _, f := range foodList {
		_, isVisited := visited[f.Name]
		if !isVisited {
			visited[f.Name] = true
			result = append(result, f.Name)
		}
	}
	log.Debugln(result)
	return result
}

func countCalories(calories, weight float64) float64 {
	return calories * weight / float64(100)
}

func rowsToMap(rows *sql.Rows) (map[string]float64, error) {
	itemsMap := map[string]float64{}
	for rows.Next() {
		var name string
		var calories float64
		err := rows.Scan(&name, &calories)
		if err != nil {
			return nil, err
		}
		itemsMap[name] = calories
	}
	return itemsMap, nil
}

func countAvailableCalories(foodList types.FoodList, caloriesMap map[string]float64) (map[string]float64, []string) {
	var missingDataFoods []string
	countedCalories := map[string]float64{}
	for _, f := range foodList {
		if _, found := caloriesMap[f.Name]; found {
			if _, foodDuplicate := caloriesMap[f.Name]; foodDuplicate {
				countedCalories[f.Name] += countCalories(caloriesMap[f.Name], float64(f.Weight))
			} else {
				countedCalories[f.Name] = countCalories(caloriesMap[f.Name], float64(f.Weight))
			}
		} else {
			missingDataFoods = append(missingDataFoods, f.Name)
		}
	}
	return countedCalories, missingDataFoods
}

func ComputeCalories(foodList types.FoodList) (float64, error) {

	rows, err := db.Connection.Query(
		"SELECT DISTINCT name,kilocalories FROM foodschema.food WHERE name = ANY($1);",
		pq.Array(getDistinctNames(foodList)))
	if err != nil {
		return -1, err
	}

	foundCaloriesMap, err := rowsToMap(rows)
	if err != nil {
		return -1, err
	}

	countedCalories, missingDataFoods := countAvailableCalories(foodList, foundCaloriesMap)

	var caloriesSum float64
	for _, c := range countedCalories {
		caloriesSum += c
	}

	log.Println(missingDataFoods)
	log.Println(countedCalories)
	log.Println(caloriesSum)

	if len(missingDataFoods) == 0 {
		return caloriesSum, nil
	}

	return caloriesSum, fmt.Errorf("missing data for some of inputed foods")
}

func buildResponse(err error, sum float64, missingDataList []string, countedCaloriesMap map[string]float64) ([]byte, error) {
	map1 := map[string]any{
		"errors":         err,
		"sum":            sum,
		"missingDataFor": missingDataList,
		"details":        countedCaloriesMap,
	}
	return json.Marshal(map1)
}
