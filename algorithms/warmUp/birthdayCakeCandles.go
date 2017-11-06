package main

import "fmt"

func main() {
	var count, temp uint64
	cand := uint64(0)
	max := uint64(0)
	fmt.Scanf("%v", &count)
	for i := uint64(0); i < count; i++ {
		fmt.Scanf("%v", &temp)
		if temp == max {
			cand++
		} else if temp > max {
			cand = 1
			max = temp
		}
	}
	fmt.Printf("%v", cand)
}
