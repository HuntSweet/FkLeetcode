package main

func twoSum(nums []int, target int) []int {
	//思路：遍历数组中的每个数，并存入字典中，边遍历边查找字典中是否有target-nums[i]的值，如果有返回下标
	m := make(map[int]int,0)
	res := []int{}
	for k,v := range nums{
		if mv,ok := m[target-v];ok{
			res = append(res,mv)
			res = append(res,k)
			return res
		}else{
			m[v] = k
		}
	}
	return res
}
