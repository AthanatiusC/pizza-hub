package helper

import (
	"math/rand"

	"github.com/AthanatiusC/pizza-hub/model"
)

func GenerateRandomUniqueID() string {
	uniqueID := make([]rune, 8)
	for i := range uniqueID {
		uniqueID[i] = model.LetterRunes[rand.Intn(len(model.LetterRunes))]
	}
	return string(uniqueID)
}
