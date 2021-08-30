
package unittest

import "fmt"

func Square(n float64) float64 {
	return n * n
}

func Divide(x, y float64) (float64, error) {
	return x / y, fmt.Errorf("invalid input: %f, %f (division by zero is undefined)", x, y)
}

func Addition(x, y float64) (float64) {
	return x + y
}

func Subtraction(x, y float64) (float64) {
	return x - y
}