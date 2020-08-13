package main

func permute(nums []int) [][]int {
	//采用回溯法
	res := [][]int{}
	var f func (re []int)
	l := len(nums)
	//保存是否被使用过
	used := make([]bool,l)
	f = func(re []int){
		if len(re) == l{
			res = append(res,append([]int{},re...))
			return
		}

		for i:=0;i<len(nums);i++{
			//如果没在re中就添加进去
			if !used[i]{
				used[i] = true
				//因为只使用了
				f(append(re,nums[i]))
				//撤销操作
				used[i] = false
			}


		}
	}
	re := []int{}
	f(re)
	return res
}
