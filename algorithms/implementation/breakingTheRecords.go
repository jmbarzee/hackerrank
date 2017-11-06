package main

import "fmt"

func main() {
	var count, min, max int64
	minB := int64(0)
	maxB := int64(0)

	fmt.Scanf("%v", &count)
	scores := make([]int64, count)

	for i := int64(0); i < count; i++ {
		fmt.Scanf("%v", &scores[i])
	}
	if count > 0 {
		var first int64
		first, scores = scores[0], scores[1:]
		min = first
		max = first
		for _, score := range scores {
			if score > max {
				max = score
				maxB++
			} else if score < min {
				min = score
				minB++
			}
		}
	}
	fmt.Printf("%v %v", maxB, minB)

}
