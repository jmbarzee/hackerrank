package main

import "fmt"

func main() {
	var num, goal, i int
	var arr []int

	fmt.Scanf("%v", &num)
	arr = make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%v", &arr[i])
	}
	goal = arr[num-1]

	for i = num - 2; i >= 0; i-- {
		if arr[i] > goal {
			arr[i+1] = arr[i]
		} else {
			arr[i+1] = goal
			break
		}
		printA(arr)
	}
	if i == -1 {
		arr[0] = goal
	}
	printA(arr)
}

func printA(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println()
}
