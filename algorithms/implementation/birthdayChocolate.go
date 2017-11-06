package main

import "fmt"

func main() {
	var count, month, day int64
	combos := int64(0)

	fmt.Scanf("%v", &count)
	nums := make([]int64, count)

	for i := int64(0); i < count; i++ {
		fmt.Scanf("%v", &nums[i])
	}
	fmt.Scanf("%v %v", &day, &month)

	if month > count {
		fmt.Printf("0")
	} else {
		sum := int64(0)
		for i := int64(0); i < month; i++ {
			sum += nums[i]
			//fmt.Printf("sum:%v nums[%v]:%v\n", sum, i, nums[i])
		}
		last := int64(0)
		next := month
		for next <= int64(len(nums)) {
			if sum == day {
				combos++
			}
			if next == int64(len(nums)) {
				break
			}
			//fmt.Printf("sum:%v nums[%v]:%v nums[%v]:%v\n", sum, last, nums[last], next, nums[next])
			sum -= nums[last]
			sum += nums[next]
			last++
			next++
		}
		fmt.Printf("%v", combos)
	}
}
