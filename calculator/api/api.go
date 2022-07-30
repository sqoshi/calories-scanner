package api

import (
	"calculator/computer"
	"calculator/translator"
	"calculator/types"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"sync"
)

// GetEnvOrFallback GetEnv but with default value
func GetEnvOrFallback(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

type ComputeResult struct {
	CaloriesSum  float64            `json:"CaloriesSum"`
	CaloriesList map[string]float64 `json:"CaloriesList"`
	Incomplete   bool               `json:"Incomplete"`
}

type Response struct {
	Error          error          `json:"Error"`
	MissingDataFor []string       `json:"MissingDataFor"`
	Result         *ComputeResult `json:"Result"`
}

func computeCaloriesSum(caloriesList map[string]float64) float64 {
	var caloriesSum float64
	for _, c := range caloriesList {
		caloriesSum += c
	}
	return caloriesSum
}

func NewResponse(err error, missingItems []string, caloriesList map[string]float64) Response {
	log.Debugln("caloriesList", caloriesList)
	return Response{err, missingItems, &ComputeResult{computeCaloriesSum(caloriesList), caloriesList, len(missingItems) != 0}}
}

func HandleCaloriesRequest(ctx *gin.Context) {
	var foodList types.FoodList
	err := json.NewDecoder(ctx.Request.Body).Decode(&foodList)
	log.Debugln("Request food list: ", foodList)
	translator.TranslateFoodNamesToEnglish(foodList)
	if err == nil {
		availableComputedCalories, missingDataFoods, dbErr := computer.ComputeCalories(foodList)
		if dbErr != nil {
			if dbErr == computer.MissingDataError {
				ctx.IndentedJSON(http.StatusOK, NewResponse(dbErr, missingDataFoods, availableComputedCalories))
				return
			}
			ctx.IndentedJSON(http.StatusOK, NewResponse(err, missingDataFoods, availableComputedCalories))
			return
		}
		ctx.IndentedJSON(http.StatusOK, NewResponse(err, missingDataFoods, availableComputedCalories))
		return
	}
	ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Could not decode request."})
}

// RunAPI deploys api with endpoints
func RunAPI(wg *sync.WaitGroup) {
	defer wg.Done()
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/compute-calories", HandleCaloriesRequest)
	addr := fmt.Sprintf("%s:%s",
		GetEnvOrFallback("DEPLOY_HOST", "0.0.0.0"),
		GetEnvOrFallback("DEPLOY_PORT", "8080"))
	err := router.Run(addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Successfully deployed on %s", addr)
}
