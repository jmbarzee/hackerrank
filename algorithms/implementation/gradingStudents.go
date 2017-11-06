package main

import "fmt"

func main() {
	var count, grade, rem int64
	fmt.Scanf("%v", &count)
	for i := int64(0); i < count; i++ {
		fmt.Scanf("%v", &grade)
		rem = grade % 5
		if grade >= 38 && rem >= 3 {
			grade += 5 - rem
		}
		fmt.Println(grade)
	}
}
