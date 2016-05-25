package projecteuler

import (
	"fmt"
)

func do() {
	sum := 0
	low := 1
	high := 2
	next := 3
	for high < 4000000 {
		if high%2 == 0 {
			sum += high
		}
		fmt.Printf("low: %d, high: %d, next: %d\n", low, high, next)
		low = high
		high = next
		next = low + high
	}
	fmt.Printf("Sum: %d\n", sum)
}
