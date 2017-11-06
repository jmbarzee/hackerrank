package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
)

func main() {
	f, _ := os.Create("profile")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var num, tmp int
	var txt string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	num, _ = strconv.Atoi(scanner.Text())
	arr := make([][]string, 100)
	for i := 0; i < 100; i++ {
		arr[i] = make([]string, 0, 100)
	}
	for i := 0; i < num; i++ {
		scanner.Scan()
		tmp, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		txt = scanner.Text()
		//fmt.Scanf("%v %v", &tmp, &txt)
		if i*2 < num {
			txt = "-"
		}
		arr[tmp] = append(arr[tmp], txt)
	}
	for i := 0; i < 100; i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%s ", arr[i][j])
		}
	}
}
