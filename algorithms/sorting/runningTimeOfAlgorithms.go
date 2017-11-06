package main

import "fmt"

func main() {
	var num, goal int
	var arr []int

	fmt.Scanf("%v", &num)
	arr = make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%v", &arr[i])
	}

	var i int
	shifts := 0

	for r := 1; r < num; r++ { // right marker, exclusive
		goal = arr[r]
		for i = r - 1; i > -1; i-- {
			if arr[i] > goal {
				arr[i+1] = arr[i]
				shifts++
			} else {
				arr[i+1] = goal
				break
			}
		}
		if i == -1 {
			arr[0] = goal
		}
	}
	fmt.Println(shifts)
}

func printA(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println()
}
