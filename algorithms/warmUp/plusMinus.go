package main

import "fmt"

func main() {
	var count, temp int32
	var pos, neg, zer int32
	fmt.Scanf("%v", &count)

	for i := int32(0); i < count; i++ {
		fmt.Scanf("%v", &temp)
		if temp > 0 {
			pos++
		} else if temp < 0 {
			neg++
		} else {
			zer++
		}
	}
	fmt.Printf("%08.6v\n", float64(pos)/float64(count))
	fmt.Printf("%08.6v\n", float64(neg)/float64(count))
	fmt.Printf("%08.6v\n", float64(zer)/float64(count))
}
