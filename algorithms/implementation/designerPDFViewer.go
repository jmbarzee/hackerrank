package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 26)
	for i := 0; i < 26; i++ {
		fmt.Scanf("%v", &arr[i])
	}

	max := 0
	var s []byte
	fmt.Scanf("%v", &s)

	for _, b := range s {
		if arr[b-'a'] > max {
			max = arr[b-'a']
		}
	}
	fmt.Println(max * len(s))
}
