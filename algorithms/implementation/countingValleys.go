package main

import "fmt"

func main() {
	var n, vals, elv int
	fmt.Scanf("%v", &n)
	steps := make([]byte, n)
	fmt.Scanf("%v", &steps)
	for _, step := range steps {
		if step == 'U' {
			elv++
			if elv == 0 {
				vals++
			}
		} else {
			elv--
		}
	}
	fmt.Printf("%v", vals)
}
