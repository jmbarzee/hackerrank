package main

import "fmt"

func main() {
	var n, birdType int

	fmt.Scanf("%v", &n)

	birdCounts := make([]int, 5)

	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &birdType)
		birdCounts[birdType-1]++
	}

	win := 0

	for j := 0; j < 5; j++ {
		if birdCounts[j] > birdCounts[win] {
			win = j
		}
	}
	fmt.Println(win + 1)
}
