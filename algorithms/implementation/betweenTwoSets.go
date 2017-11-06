package main

import "fmt"

func main() {
	var n, m int64
	fmt.Scanf("%v %v", &n, &m)

	nSet := make([]int64, n)
	mSet := make([]int64, m)
	nMax := int64(0)
	mMin := int64(101)

	for i := int64(0); i < n; i++ {
		fmt.Scanf("%v", &nSet[i])
		if nMax < nSet[i] {
			nMax = nSet[i]
		}
	}
	for i := int64(0); i < m; i++ {
		fmt.Scanf("%v", &mSet[i])
		if mMin < mSet[i] {
			mMin = mSet[i]
		}
	}

	posx := make([]int64, 0)
	for i := nMax; i <= mMin; i++ {
		valid := true // true
		for _, m := range mSet {
			if m%i != 0 {
				valid = false
				break //optomization
			}
		}
		for _, n := range nSet {
			if i%n != 0 {
				valid = false
				break
			}
		}
		if valid {
			posx = append(posx, i)
		}
	}
	fmt.Printf("%v", len(posx))
}
