package api

import (
	"calculator/computer"
	"calculator/types"
	"encoding/json"
	"fmt"
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

func getDataFromRequest(ctx *gin.Context) {
	var foodList types.FoodList
	err := json.NewDecoder(ctx.Request.Body).Decode(&foodList)
	if err == nil {
		log.Println(foodList)
		calories, dbErr := computer.ComputeCalories(foodList)
		if dbErr != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": dbErr})
			return
		}
		ctx.IndentedJSON(http.StatusOK, calories)
		return
	}
	ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Could not decode request."})
}

// RunAPI deploys api with endpoints
func RunAPI(wg *sync.WaitGroup) {
	defer wg.Done()
	router := gin.Default()
	router.POST("/compute-calories", getDataFromRequest)
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
