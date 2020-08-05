package dynamicPlan

func integerBreak(n int) int {
	dp := make([]int,n+1)
	dp[1] = 1
	dp[2] = 1

	for i:=3;i<=n;i++{
		for j:=1;j<=n-1;j++{
			dp[i] = max(dp[j]*(i-j),dp[i])
			dp[i] = max(j*(i-j),dp[i])
		}
	}
	return dp[n]
}

func max(x,y int) int{
	if x > y{
		return x
	}

	return y
}