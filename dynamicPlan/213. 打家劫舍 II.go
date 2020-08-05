package dynamicPlan

func rob(nums []int) int {
	//思路：因为首尾会报警，排除首尾分别计算最大值之后取它们之间的最大值
	n := len(nums)

	if n <= 0{
		return 0
	}
	//记住nums[:0]会等于[]
	if n == 1{
		return nums[0]
	}
	res := max(robs(nums[:n-1]),robs(nums[1:]))
	return res
}

func robs(nums []int) int{
	pre,cur := 0,0
	for _,v := range nums{
		cur,pre = max(v+pre,cur),cur
	}
	return cur
}
func max(x,y int) int{
	if x > y{
		return x
	}

	return y
}

