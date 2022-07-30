package translator

import (
	"calculator/types"
	"fmt"
	gt "github.com/bas24/googletranslatefree"
	"strings"
)

func TranslateFoodNamesToEnglish(foods []types.Food) {
	for i := range foods {
		translatedName, err := gt.Translate(foods[i].Name, "auto", "en")
		fmt.Println(err)
		if err == nil {
			foods[i].Name = strings.ToLower(translatedName)
		}
	}
	fmt.Println(foods)
}

///func TranslateFoodNamesToEnglish(foods []*types.Food) {
//	for _, food := range foods {
//		translatedName, err := gt.Translate(food.Name, "auto", "en")
//		fmt.Println(err)
//		if err == nil {
//			food.Name = strings.ToLower(translatedName)
//		}
//	}
//	fmt.Println(&foods)
//}
