package projecteuler

import (
	"fmt"
)

func doS() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}