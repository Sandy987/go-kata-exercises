// Attempted solution for https://www.codewars.com/kata/57339a5226196a7f90001bcf

package kata

import (
	"sort"
)

// Crossover  calculates an n-point genetic crossover of two chromosomes
// ns : slice of indices
// xs, ys : chromosomes as slices of ints
func Crossover(ns []int, xs []int, ys []int) ([]int, []int) {
	newX := make([]int, len(xs))
	newY := make([]int, len(ys))
	crossOvers := sortAndRemoveDupes(ns)
	isCrossed := false
	crossIndex := 0
	for i := 0; i < len(xs); i++ {
		if crossIndex < len(crossOvers) {
			crossPoint := crossOvers[crossIndex]
			if crossPoint == i {
				isCrossed = !isCrossed
				crossIndex++
			}
		}

		if isCrossed {
			newX[i] = ys[i]
			newY[i] = xs[i]
		} else {
			newX[i] = xs[i]
			newY[i] = ys[i]
		}
	}
	return newX, newY
}

func sortAndRemoveDupes(arr []int) []int {
	vals := make(map[int]bool)
	var outarr []int
	for _, c := range arr {
		vals[c] = true
	}
	for k := range vals {
		outarr = append(outarr, k)
	}
	sort.Ints(outarr)
	return outarr
}
