package main

import "sort"

func permuteUnique(nums []int) [][]int {
	//思路：先排序，回溯加剪枝
	// 排序（升序或者降序都可以），排序是剪枝的前提
	//https://leetcode-cn.com/problems/permutations-ii/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liwe-2/
	l := len(nums)
	res := [][]int{}
	if l == 0{
		return res
	}
	sort.Ints(nums)
	used := make([]bool,l)
	var f func(re []int)

	f = func(re []int) {
		if len(re) == l {
			//注意这里防止被覆盖
			res = append(res,append([]int{},re...))
			return
		}

		for i:=0;i<l;i++{
			if used[i]{
				continue
			}

			//i>0是为了让 i-1 有意义
			//used[i-1]代表刚刚被撤销选择
			if i > 0 && nums[i] == nums[i-1] && !used[i-1]{
				continue
			}
			//选择
			used[i] = true

			f(append(re,nums[i]))
			//撤销选择

			used[i] = false
		}
	}

	f([]int{})
	return res
}
