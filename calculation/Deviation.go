package calculation

import "math"

// Deviation calculates the standard deviation of a slice of float64 numbers.
func Deviation(data []float64) float64 {
	return math.Sqrt(Variance(data))
}
