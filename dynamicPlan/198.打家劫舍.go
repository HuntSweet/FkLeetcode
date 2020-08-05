package dynamicPlan

//状态：dp[i]为最大金额
//要么dp[i+1]=dp[i-1]+nums[i],要么等于dp[i]
func rob(nums []int) int {
	//因为只依赖于两个值，所以直接用两个变量动态更新，pre=dp[i-1],cur=dp[i]
	//pre,cur := 0,0
	//for _,v := range nums{
	//	cur,pre = max(v+pre,cur),cur
	//}
	//return cur
	n := len(nums)
	if n <= 0{
		return 0
	}
	dp := make([]int,n+1)
	dp[1] = nums[0]
	for i:=1;i<len(nums);i++{
		dp[i+1] = max(dp[i-1]+nums[i],dp[i])
	}

	return dp[n]
}

func max(x,y int) int{
	if x > y{
		return x
	}

	return y
}