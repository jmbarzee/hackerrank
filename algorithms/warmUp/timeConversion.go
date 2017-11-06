package main

import "fmt"

func main() {
	var hor, min, sec int
	var half string
	fmt.Scanf("%v:%v:%v%v", &hor, &min, &sec, &half)
	if hor == 12 {
		hor = 0
	}
	if half == "PM" {
		hor += 12
	}
	fmt.Printf("%02v:%02v:%02v", hor, min, sec)
}
