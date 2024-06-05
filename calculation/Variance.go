package calculation

// Variance calculates the variance of a slice of float64 numbers.
func Variance(data []float64) float64 {
	avg := Average(data)
	sum := 0.0
	for _, value := range data {
		sum += (value - avg) * (value - avg)
	}
	return sum / float64(len(data))
}
