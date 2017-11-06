package main

import "fmt"

func main() {
	var aTree int64 // leftmost object
	var lHome int64
	var rHome int64
	var oTree int64 // rightmost object

	var aNum, aHit, oNum, oHit, relPos, absPos int64

	fmt.Scanf("%v %v", &lHome, &rHome)
	fmt.Scanf("%v %v", &aTree, &oTree)
	fmt.Scanf("%v %v", &aNum, &oNum)

	for i := int64(0); i < aNum; i++ {
		fmt.Scanf("%v", &relPos)
		absPos = aTree + relPos
		if absPos >= lHome && absPos <= rHome {
			aHit++
		}
	}

	for i := int64(0); i < oNum; i++ {
		fmt.Scanf("%v", &relPos)
		absPos = oTree + relPos
		if absPos >= lHome && absPos <= rHome {
			oHit++
		}
	}

	fmt.Println(aHit)
	fmt.Println(oHit)
}
