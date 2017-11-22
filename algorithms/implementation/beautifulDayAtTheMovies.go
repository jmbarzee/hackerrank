package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var i, j, k, n, m, c int
	fmt.Scanf("%v %v %v", &i, &j, &k)
	for n = i; n <= j; n++ {
		m = Reverse(n)
		if Abs(m-n)%k == 0 {
			c++
		}
	}
	fmt.Println(c)
}

func Reverse(i int) int {
	s := strconv.Itoa(i)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	i, _ = strconv.Atoi(string(runes))
	return i
}

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}
