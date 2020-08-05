package main

import (
	"fmt"
)
type ListNode struct {
	Val int
	Next *ListNode
}
func main(){
	fmt.Println(10)
	ld := new(ListNode)
	fmt.Println(ld.Val)

	dp := make([][]int,3)
	for i:=0;i<3;i++{
		if dp[i] == nil{
			dp[i] = make([]int,2)
		}

	}
	fmt.Println(dp)
}
