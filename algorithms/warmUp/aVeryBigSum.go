package main

import "fmt"

func main() {
	var count, temp uint32
	total := uint64(0)
	fmt.Scanf("%v", &count)
	for i := uint32(0); i < count; i++ {
		fmt.Scanf("%v", &temp)
		total += uint64(temp)
	}
	fmt.Printf("%d", total)
}
