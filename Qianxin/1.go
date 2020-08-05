package main

import "fmt"

func bestValue(s,n int,w,v []int) int {
	dp := make([]int,s+1)
	for i:=1;i<n;i++{
		for j:=w[i];j<=s;j++{
			dp[j] = max(dp[j-w[i]]+v[i],dp[j])
		}
	}

	return dp[s]
}
func max(x,y int) int  {
	if x > y{
		return x
	}

	return y
}
func main()  {
	//背包容量，物品总数
	var s,n int
	fmt.Scan(&s)
	fmt.Scan(&n)
	//容量，价值
	w := make([]int,0)
	v := make([]int,0)

	for i:=0;i<n;i++{
		var temp1,temp2 int
		fmt.Scanf("%d %d",&temp1,&temp2)
		w = append(w,temp1)
		v = append(v,temp2)
	}

	fmt.Println(bestValue(s,n,w,v))

}