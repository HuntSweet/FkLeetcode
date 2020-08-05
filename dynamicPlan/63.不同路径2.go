package dynamicPlan

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	dp := make([][]int,n)
	for i:=0;i<n;i++{
		dp[i] = make([]int,m)
		for j:=0;j<m;j++{
			if obstacleGrid[i][j] == 1{
				dp[i][j] = 0
			}else{
				dp[i][j] = 1
			}

		}
	}

	for i:=n-1;i>=0;i--{
		for j:=m-1;j>=0;j--{

			if j != m-1 && i == n-1{
				dp[i][j] *= dp[i][j+1]
			}else if j == m-1 && i!= n-1{
				dp[i][j] *= dp[i+1][j]
			}else if j != m-1 && i != n-1{
				dp[i][j] *= dp[i+1][j] + dp[i][j+1]
			}else{
				dp[i][j] = dp[i][j]
			}
		}
	}

	return dp[0][0]
}
