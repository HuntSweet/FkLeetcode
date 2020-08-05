package dynamicPlan
//动态规划

func trap(height []int) int {
	if len(height) <=2 {
		return 0
	}
	//思路：求每一列能接到的雨水后相加（最左边和最右边肯定接不到雨水），当且仅当每一列左边和右边最大值中的较小值大于当前列，才能接到雨水
	res := 0
	//为了防止重复寻找每一列的左右最大值，分别用动态规划数组存储每一列左边的最大值和右边的最大值
	dp := make([]int,len(height))
	dp[0] = height[0]
	for i:=1;i<len(dp);i++{
		dp[i] = max(dp[i-1],height[i])
	}

	dpr := make([]int,len(height))
	dpr[len(dpr)-1] = height[len(dpr)-1]
	for i:=len(dpr)-2;i>=0;i--{
		dpr[i] = max(dpr[i+1],height[i])
	}

	for i:=1;i<=len(height)-2;i++{
		left := dp[i]
		right := dpr[i]

		m := min(left,right)
		if m > height[i]{
			res += m - height[i]
		}

	}
	return res
}
func max(x,y int)int{
	if x > y{
		return x
	}
	return y
}
func min(x,y int)int{
	if x > y{
		return y
	}
	return x
}