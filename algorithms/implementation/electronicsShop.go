package main

import (
	"fmt"
	"sort"
)

// ByVal implements sort.Interface for []int
type ByVal []int

func (a ByVal) Len() int           { return len(a) }
func (a ByVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVal) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	var s, n, m int
	fmt.Scanf("%v %v %v", &s, &n, &m)
	narr := make([]int, n)
	marr := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v", &narr[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%v", &marr[i])
	}
	sort.Sort(ByVal(narr))
	sort.Sort(ByVal(marr))

	ms := -1
	j := 0
	for i := n - 1; i > -1; i-- {
		for ; j < m; j++ {
			if narr[i]+marr[j] > s {
				break
			} else if narr[i]+marr[j] > ms {
				ms = narr[i] + marr[j]
			}
		}
	}
	fmt.Println(ms)
}
