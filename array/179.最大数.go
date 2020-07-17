package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	ss := ""

	sort.Slice(nums, func(i, j int) bool {
		t1 := strconv.Itoa(nums[i])
		t2 := strconv.Itoa(nums[j])

		z1,_ := strconv.Atoi(t1+t2)
		z2,_ := strconv.Atoi(t2+t1)
		return  z1 > z2
	})
	//注意为0的情况
	if nums[0] == 0{
		return "0"
	}
	for _,v := range nums{
		ss += strconv.Itoa(v)
	}

	return ss
}

//func sort()  {
//
//}
func main()  {
	v := []int{9,30,34,5,3}
	fmt.Println(largestNumber(v))
}