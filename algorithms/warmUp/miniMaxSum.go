package main

import "fmt"
import "math"

func main() {
	var temp uint64
	min := uint64(math.MaxUint64)
	max := uint64(0)
	total := uint64(0)
	for i := 0; i < 5; i++ {
		fmt.Scanf("%v", &temp)
		total += temp
		if temp < min {
			min = temp
		}
		if temp > max {
			max = temp
		}
	}
	fmt.Printf("%v %v", total-max, total-min)
}
