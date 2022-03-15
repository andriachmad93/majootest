package main

import (
	"fmt"
	"sort"
)

var result = []int{}

func SortAscDesc(arrlist []int, asc bool) []int {
	// descending
	sort.Slice(arrlist, func(i, j int) bool {
		return arrlist[i] > arrlist[j]
	})

	if asc {
		// ascending
		sort.Slice(arrlist, func(i, j int) bool {
			return arrlist[i] < arrlist[j]
		})
	}

	for _, v := range arrlist {
		result = append(result, v)
	}

	return result
}

func main() {
	a := []int{10, 3, 4, 11, 8, -1, -10}

	fmt.Printf("%d", SortAscDesc(a, false))
}
