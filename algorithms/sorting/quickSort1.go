package main

import "fmt"

func main() {
	var num int
	var arr []int

	fmt.Scanf("%v", &num)
	arr = make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%v", &arr[i])
	}

	equal := make([]int, 1)
	less := make([]int, 0)
	more := make([]int, 0)

	pivot := arr[0]
	equal[0] = pivot

	for i := 1; i < num; i++ {
		if arr[i] < pivot {
			less = append(less, arr[i])
		} else if arr[i] > pivot {
			more = append(more, arr[i])
		} else {
			equal = append(equal, arr[i])
		}
	}

	for i := 0; i < len(less); i++ {
		fmt.Printf("%v ", less[i])
	}
	for i := 0; i < len(equal); i++ {
		fmt.Printf("%v ", equal[i])
	}
	for i := 0; i < len(more); i++ {
		fmt.Printf("%v ", more[i])
	}
}

func printA(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println()
}
