package main

import "fmt"

func main() {
	var t, n, k, b, temp int
	fmt.Scanf("%v", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%v %v", &n, &k)
		b = 0
		for j := 0; j < n; j++ {
			fmt.Scanf("%v", &temp)
			if temp <= 0 {
				b++
			}
		}
		if b >= k {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
