package main

import (
	"fmt"
	"sort"
)

type ByVal []string

func (a ByVal) Len() int      { return len(a) }
func (a ByVal) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByVal) Less(i, j int) bool {
	if len(a[i]) < len(a[j]) {
		return true
	} else if len(a[i]) > len(a[j]) {
		return false
	} else {
		for k, _ := range a[i] {
			if a[i][k] < a[j][k] {
				return true
			} else if a[i][k] > a[j][k] {
				return false
			}
		}
		return false
	}
}

func main() {
	var count uint64
	fmt.Scanf("%v", &count)
	vals := make([]string, count)
	for i := uint64(0); i < count; i++ {
		fmt.Scanf("%v", &vals[i])
	}
	sort.Sort(ByVal(vals))

	for i := uint64(0); i < count; i++ {
		fmt.Printf("%v\n", vals[i])
	}
}
