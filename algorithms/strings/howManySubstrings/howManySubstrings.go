package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sort"
)

type job struct {
	s    string
	i, j int
}

type ByVal []string

func (a ByVal) Len() int           { return len(a) }
func (a ByVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVal) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	f, _ := os.Create("profile")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var t, k, i, j int
	fmt.Scanf("%v %v", &t, &k)
	var s string
	fmt.Scanf("%v", &s)
	//fmt.Printf("s: %v\n", s)

	for n := 0; n < 100; n++ {
		//fmt.Printf("  n: %v\n", n)

		fmt.Scanf("%v %v", &i, &j)
		//fmt.Printf("  i: %v  j: %v\n", i, j)

		arr := make([]string, j-i+1)
		p := 0
		for m := i; m <= j; m++ {
			arr[p] = s[m : j+1] // need to be inclusive of j
			//fmt.Printf("    arr[p]:%v\n", arr[p])
			p++
		}

		sort.Sort(ByVal(arr))
		c := len(arr[0])
		for q := 1; q < len(arr); q++ {
			c += len(arr[q]) - lcp(arr[q], arr[q-1])
		}
		fmt.Println(c)
	}

}

func lcp(a, b string) int {
	min := Min(len(a), len(b))
	var i int
	for i = 0; i < min; i++ {
		if a[i] != b[i] {
			break
		}
	}
	return Min(i, min)
}

func Min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
