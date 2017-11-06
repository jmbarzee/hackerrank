package main

import "fmt"

func main() {
	var a [3]uint32
	var b [3]uint32
	atot := uint32(0)
	btot := uint32(0)
	for i := 0; i < 3; i++ {
		fmt.Scanf("%v", &a[i])
	}
	for j := 0; j < 3; j++ {
		fmt.Scanf("%v", &b[j])
	}
	for k := 0; k < 3; k++ {
		if a[k] > b[k] {
			atot++
		} else if a[k] < b[k] {
			btot++
		}
	}
	fmt.Printf("%v %v", atot, btot)
}
