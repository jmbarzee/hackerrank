package main

import "fmt"

func main() {
	var num, k, tmp int
	fmt.Scanf("%v %v", &num, &k)
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%v", &tmp)
		arr[i] = tmp % k
	}

	cnt := 0
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if (arr[i]+arr[j])%k == 0 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
