package translator

import (
	"calculator/types"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestShouldTranslate(t *testing.T) {
	foodlist := []types.Food{{Name: "AwOkAdO"}, {Name: "JAbłko"}}
	TranslateFoodNamesToEnglish(foodlist)

	assert.Equal(t, foodlist[0].Name, "avocado")
	assert.Equal(t, foodlist[1].Name, "apple")
}
