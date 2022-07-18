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
	foodArray := types.FoodList{
		types.Food{Name: "Avocado", Weight: 300},
		types.Food{Name: "Butter", Weight: 115},
		types.Food{Name: "Apple", Weight: 1000},
		types.Food{Name: "asd", Weight: 1000},
		types.Food{Name: "Apple", Weight: 1000},
		types.Food{Name: "Apple", Weight: 1000},
	}
	postBody, _ := json.Marshal(foodArray)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:8080/compute-calories", "application/json", responseBody)
	if err != nil {
		t.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
