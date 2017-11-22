package main

import "fmt"

func main() {
	var year int
	fmt.Scanf("%v", &year)

	day := 13
	if year < 1918 {
		if year%4 == 0 {
			day--
		}
		fmt.Printf("%v.09.%v", day, year)
	} else if year > 1918 {
		if (year%4) == 0 && (year%100) != 0 {
			day--
		} else if (year % 400) == 0 {
			day--
		}
		fmt.Printf("%v.09.%v", day, year)
	} else {
		fmt.Println("26.09.1918")
	}
}
