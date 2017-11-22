package main

import "fmt"

func main() {
	var pages, page, f, b int
	fmt.Scanf("%v\n%v", &pages, &page)
	if pages%2 == 0 {
		pages++
	}
	f = page / 2
	b = (pages - page) / 2
	if f > b {
		fmt.Println(b)
	} else {
		fmt.Println(f)
	}
}
