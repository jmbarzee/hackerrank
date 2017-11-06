package main

import (
	"fmt"
	"math"
)

func main() {
	var side uint32
	fmt.Scanf("%v", &side)
	var matrix = make([]int32, side*side)

	for i := uint32(0); i < side; i++ {
		for j := uint32(0); j < side; j++ {
			fmt.Scanf("%v", &matrix[i*side+j])
		}
	}

	forward := int32(0)
	backward := int32(0)

	for i := uint32(0); i < side; i++ {
		forward += matrix[i*side+i]
	}

	for i := uint32(1); i <= side; i++ {
		backward += matrix[(side-1)*i]
	}
	total := math.Abs(float64(forward - backward))
	fmt.Printf("%v", total)
}
