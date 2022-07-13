package api

import (
	"bytes"
	"calculator/types"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestSendTestRequest(t *testing.T) {
	f1 := types.Food{Name: "Avocado", Weight: 300}
	f2 := types.Food{Name: "Butter", Weight: 115}
	foodArray := types.FoodList{f1, f2}
	postBody, _ := json.Marshal(foodArray)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://0.0.0.0:8080/compute-calories", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
