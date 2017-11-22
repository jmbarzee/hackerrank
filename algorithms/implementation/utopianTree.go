package main

import "fmt"

func main() {
	var t, n uint
	fmt.Scanf("%v", &t)
	for i := uint(0); i < t; i++ {
		fmt.Scanf("%v", &n)
		o := uint(1)
		/*
		   y := n>>o
		   fmt.Printf("y: %b\n", y)
		   p := ^o<<(n>>o)
		   fmt.Printf("p: %b\n", p)
		   v := ^(^o<<(n>>o))
		   fmt.Printf("v: %b\n", v)
		*/
		s := (^(^o << (n >> o))) << (n % 2)
		fmt.Printf("%v\n", s)
	}
}

/*
 N  Y  H
 0  0  1
 1  0  2
 2  1  3
 3  1  6
 4  2  7
 5  2  14
 6  3  15
 7  3  30
 8  4  31
*/
