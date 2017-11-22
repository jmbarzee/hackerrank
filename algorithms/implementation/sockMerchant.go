package main

import "fmt"

func main() {
	var n, sock, pairs int
	fmt.Scanf("%v", &n)
	arr := make([]int, 100)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &sock)
		arr[sock-1]++
	}
	for i := 0; i < 100; i++ {
		pairs += arr[i] / 2
	}
	fmt.Println(pairs)
}
