package main

import "fmt"

func main() {
	var n, skip, total, skipCost, charge, temp int64
	fmt.Scanf("%v %v", &n, &skip)
	for i := int64(0); i < n; i++ {
		fmt.Scanf("%v", &temp)
		total += temp
		if i == skip {
			skipCost = temp
		}
	}
	fmt.Scanf("%v", &charge)

	if charge*2 == total-skipCost {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Printf("%v", (charge - (total-skipCost)/2))
	}
}
