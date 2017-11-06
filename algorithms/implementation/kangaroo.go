package main

import "fmt"

func main() {
	var x1, j1, x2, j2 float64
	fmt.Scanf("%v %v %v %v", &x1, &j1, &x2, &j2)
	inter := (x2 - x1) / (j1 - j2)
	if inter > 0 && inter == float64(int64(inter)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
