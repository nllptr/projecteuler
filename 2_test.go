package projecteuler

import (
	"testing"
)

var testData = []struct {
	intput   int
	expected int
}{
	{90, 44},
}

func TestSumFib(t *testing.T) {
	for _, td := range testData {
		actual := SumFib(td.intput)
		if actual != td.expected {
			t.Errorf("SumFir(%d): expected %d, actual %d", td.intput, td.expected, actual)
		}
	}
}
