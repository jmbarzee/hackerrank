package main

import "fmt"

func main() {
	var a, read, res uint32
	res = 0
	fmt.Scanf("%v\n", &a)
	var i uint32
	for i = 0; i < a; i++ {
		fmt.Scanf("%v", &read)
		res += read
	}
	fmt.Println(res)
}
