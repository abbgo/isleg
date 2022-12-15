package pkg

import (
	"math"
	"os"
)

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

var ServerPath = os.Getenv("SERVER_PATH")
