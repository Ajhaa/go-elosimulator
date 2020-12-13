package simulator

import (
	"math"
	"math/rand"
)

// Predict a result based on two ratings
func Predict(elo1 float64, elo2 float64) float64 {
	return 1 / (1 + math.Pow(10, (elo2-elo1)/400))
}

var selection = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm"

func randomString(n int) string {
	s := ""

	for i := 0; i < n; i++ {
		s += string(selection[rand.Intn(len(selection))])
	}

	return s
}
