// Package randsample implements the CacheDiff Random Index Selection algorithm
/*

http://arxiv.org/pdf/1512.00501v1.pdf

*/
package randsample

import (
	"math/rand"
)

func Sample(N, k int) []int {

	me := make(map[int]int)
	output := make([]int, 0, k)

	for i := N - 1; i >= N-k; i-- {
		j := rand.Intn(i + 1)
		var indexJ, indexI int
		if meJ, ok := me[j]; ok {
			indexJ = meJ
		} else {
			indexJ = j
		}

		if meI, ok := me[i]; ok {
			indexI = meI
		} else {
			indexI = i
		}
		me[i] = indexI
		me[j] = indexJ
		output = append(output, indexJ)
	}

	return output
}
