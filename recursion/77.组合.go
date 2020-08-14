package main

func combine(n int, k int) [][]int {
	res := [][]int{}
	if n == 0{
		return res
	}

	var f func(index int,re []int)
	//index表示不考虑前面已经选取的数，避免重复
	f = func(index int,re []int){
		if len(re) == k{
			res = append(res,append([]int{},re...))
			return
		}

		//[i,n]分别区间两端分别剪枝
		for i:=index;i<=n-(k-len(re))+1;i++{
			f(i+1,append(re,i))
		}
	}

	f(1,[]int{})
	return res
}