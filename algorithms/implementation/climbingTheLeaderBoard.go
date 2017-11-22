package main

import (
	"fmt"
)

func main() {
	var n, m, temp int
	fmt.Scanf("%v", &n)
	ranks := make([]int, 1)    // don't make array size n
	fmt.Scanf("%v", &ranks[0]) // this makes our loop cleaner, 1 <= n
	for i := 1; i < n; i++ {

		fmt.Scanf("%v", &temp)
		// this serves the purpose of only recording scores
		// that add to the rank.
		// this takes the solution from O(n) to O(m)
		// where m is a unique set of elements in n
		if temp < ranks[len(ranks)-1] {
			ranks = append(ranks, temp)
		}
	}
	// ranks is now a map of rank-1 => score required for rank

	fmt.Scanf("%v", &m)
	scores := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%v", &scores[i])
	}

	// r actually represents our rank, easy off by one errors ahead
	r := len(ranks) + 1
	for _, s := range scores {

		// raise our rank until we are first or below a higher scoring rank
		for r > 1 && ranks[r-2] <= s {
			r--
		}
		fmt.Println(r)
	}
}
