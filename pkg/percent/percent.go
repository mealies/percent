package percent

import (
	"github.com/mealies/percent/internal/resources"
	"golang.org/x/exp/constraints"
	"math"
)

const (
	PercentMin = 0.0
	PercentMax = 100.0
)

// Percent returns the percentage of value.
func Percent[T constraints.Integer | constraints.Float](percent, value T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resources.ErrOutOfRange
	}

	return float64(value) * (float64(percent) / PercentMax), nil
}

// Of calculates the percentage of the part relative to the total.
func Of[T constraints.Integer | constraints.Float](part, total T) (float64, error) {
	if float64(total) == 0 {
		return 0, resources.ErrDivideByZero
	}

	if float64(part) > float64(total) {
		return 0, resources.ErrPartGreaterThanTotal
	}

	return float64(part) / float64(total) * PercentMax, nil
}

// Change calculates the percentage change between two values.
func Change[T constraints.Integer | constraints.Float](oldValue, newValue T) (float64, error) {
	if float64(oldValue) == 0 {
		return 0, resources.ErrDivideByZero
	}

	return (float64(newValue) - float64(oldValue)) / math.Abs(float64(oldValue)) * PercentMax, nil
}

// Remain returns the percentage of value that remains after subtracting the percentage.
func Remain[T constraints.Integer | constraints.Float](percent, value T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resources.ErrOutOfRange
	}

	return float64(value) * ((PercentMax - float64(percent)) / PercentMax), nil
}

// FromRatio returns the percent of ratio.
func FromRatio[T constraints.Integer | constraints.Float](ratio T) (float64, error) {
	if float64(ratio) < 0 || float64(ratio) > 1 {
		return 0, resources.ErrOutOfRange
	}

	return float64(ratio) * PercentMax, nil
}

// ToRatio returns the ratio of percent.
func ToRatio[T constraints.Integer | constraints.Float](percent T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resources.ErrOutOfRange
	}

	return float64(percent) / PercentMax, nil
}
