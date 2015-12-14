package randsample

import (
	"testing"
)

func TestSample(t *testing.T) {

	var counts [100]int

	for i := 0; i < 1e6; i++ {
		for _, v := range Sample(100, 10) {
			counts[v]++
		}
	}

	for i, v := range counts {
		t.Logf("%02d %d\n", i, v)
	}
}
