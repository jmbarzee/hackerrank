package main

import (
	"fmt"
	"math"
)

func main() {
	var q, a, b, c int
	fmt.Scanf("%v", &q)
	for i := 0; i < q; i++ {
		fmt.Scanf("%v %v %v", &a, &b, &c)
		if abs(a-c) > abs(b-c) {
			fmt.Println("Cat B")
		} else if abs(a-c) < abs(b-c) {
			fmt.Println("Cat A")
		} else {
			fmt.Println("Mouse C")
		}
	}
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
