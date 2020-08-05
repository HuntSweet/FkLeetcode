package dynamicPlan


func numSquares2(n int) int {
	dp := make([]int,n+1)

	dp[0] = 0
	dp[1] = 1
	for i:=2;i<=n;i++{
		dp[i] = i
		for j:=1;i - j*j >=0;j++{
			dp[i] = min(dp[i-j*j]+1,dp[i])
		}

	}

	return dp[n]
}

func numSquares(n int) int {
	dp := make([]int,n+1)
	lt := make([]int,0)
	for i := 1;i*i<=n;i++{
		lt = append(lt,i*i)
	}

	dp[0] = 0
	dp[1] = 1
	for i:=2;i<=n;i++{
		dp[i] = i
		for j:=0;j < len(lt) && i>=lt[j];j++{

			dp[i] = min(dp[i-lt[j]]+1,dp[i])


		}

	}

	return dp[n]
}

func min(x,y int) int{
	if x > y{
		return y
	}

	return x
}
