package types

type Food struct {
	Name   string `json:"Name"`
	Weight int    `json:"Weight"`
}
type FoodList []Food
