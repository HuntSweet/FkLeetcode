package main

import "fmt"

/**
 *
 * @param N int整型
 * @return int整型
 */
//1,4,16,64
func GetCoinCount( N int ) int {
	// write code here
	n := 1024-N
	dp := make([]int,n+1)
	for k,_ := range dp{
		dp[k] = k
	}
	for k,_ := range dp{

		if k == 4{
			dp[k] = 1
		}else if k>4 && k< 16{
			dp[k] = min(dp[k],dp[k-4]+1)
		}else if k==16{
			dp[k] = 1
		}else if k>16 && k<64{
			dp[k] = min(dp[k-16]+1,dp[k])
		}else if k==64{
			dp[k] = 1
		} else if k<4{
			dp[k] = k
		}else{
			dp[k] = min(dp[k-64]+1,dp[k])
		}

	}

	return dp[n]

}

func min(x,y int) int {
	if x > y{
		return y
	}
	return x
}

func main()  {
	fmt.Println(GetCoinCount(0))
}