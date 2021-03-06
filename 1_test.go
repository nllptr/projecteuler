package projecteuler

import (
	"fmt"
	"testing"
)

var testData1 = []struct {
	input    int
	expected int
}{
	{10, 23},
}

func TestDoSum(t *testing.T) {
	for _, td := range testData1 {
		actual := DoSum(td.input)
		if actual != td.expected {
			t.Errorf("DoSum(%d): expected %d, actual %d", td.input, td.expected, actual)
		}
	}
	fmt.Printf("1: DoSum(1000) = %d\n", DoSum(1000))
}
