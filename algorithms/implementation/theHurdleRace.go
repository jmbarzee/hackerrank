package main

import (
	"fmt"
)

func main() {
	var n, k, hur int
	max := 0
	fmt.Scanf("%v %v", &n, &k)
	for i := 1; i < n; i++ {
		fmt.Scanf("%v", &hur)
		if hur > max {
			max = hur
		}
	}
	if k > max {
		fmt.Println(0)
	} else {
		fmt.Println(max - k)
	}
}
