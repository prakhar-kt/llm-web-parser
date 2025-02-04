
package main

import (
	"fmt"
	"time"
)

func calculate(iterations, param1, param2 int) float64 {
	result := 1.0
	for i := 1; i <= iterations; i++ {
		j := float64(i*param1 - param2)
		result -= (1 / j)
		j = float64(i*param1 + param2)
		result += (1 / j)
	}
	return result
}

func main() {
	start := time.Now()
	result := calculate(100_000_000, 4, 1) * 4
	elapsed := time.Since(start)

	fmt.Printf("Result: %.12f\n", result)
	fmt.Printf("Execution Time: %.6f seconds\n", elapsed.Seconds())
}
