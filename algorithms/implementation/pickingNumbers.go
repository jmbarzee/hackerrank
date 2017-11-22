package main

import "fmt"

func main() {
	var n, temp int
	fmt.Scanf("%v", &n)
	arr := make([]int, 100)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &temp)
		arr[temp]++
	}
	max := 0
	for i := 0; i < 99; i++ {
		if arr[i]+arr[i+1] > max {
			max = arr[i] + arr[i+1]
		}
	}
	fmt.Println(max)
}
