package percent

// Percent - calculate what %[number1] of [number2] is.
// ex. 25% of 100 is 25
func Percent(percent int, n int) float64 {
	return ((float64(n) * float64(percent)) / float64(100))
}
