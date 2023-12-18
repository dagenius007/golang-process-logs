package utils

import (
	"math"
)

func FormatTo2Decimal(value float64) float64 {
	return math.Ceil((value)*100) / 100
}
