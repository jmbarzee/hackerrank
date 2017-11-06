package main

import "fmt"

func main() {
	var count, goal int
	var l, c, r int // bounds for searching, left, center, right
	var arr []int

	fmt.Scanf("%v", &goal)
	fmt.Scanf("%v", &count)

	// fill arr with sorted list
	arr = make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scanf("%v", &arr[i])
	}

	// search for goal
	l = 0
	r = count
	for {
		c = (r-l)/2 + l
		if arr[c] < goal {
			l = c
		} else if arr[c] > goal {
			r = c
		} else {
			break
		}
	}
	fmt.Println(c)
}
