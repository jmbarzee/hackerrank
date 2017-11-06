package main

import "fmt"

func main() {
	var count int64
	fmt.Scanf("%v", &count)

	for i := int64(1); i <= count; i++ {
		for j := int64(0); j < count-i; j++ {
			fmt.Printf(" ")
		}
		for j := int64(0); j < i; j++ {
			fmt.Printf("#")
		}
		fmt.Println()
	}
}
