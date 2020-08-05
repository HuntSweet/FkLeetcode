package main

import "sort"

func threeSum(nums []int) [][]int {
	//排序+双指针
	//遍历，如果大于0，说明后面都大于0，无解，跳过，如果加起来等于0，那么加入结果，如果大于0，那么r--，如果小于0，l++
	//避免重复解：如果有相等元素，那么就跳过，如果得到解，那么继续判断是否相等，如果相等，继续跳过
	if len(nums) < 3{
		return nil
	}
	sort.Ints(nums)
	res := make([][]int,0)
	for i := 0;i<len(nums)-2;i++{
		if nums[i] > 0{
			return res
		}
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}

		l := i+1
		r := len(nums)-1

		for l < r{
			if nums[r] + nums[l] + nums[i] == 0{
				temp := []int{nums[i],nums[l],nums[r]}
				res = append(res,temp)
				for l<r && nums[l+1] == nums[l]{
					l++
				}
				for l < r && nums[r-1] == nums[r]{
					r--
				}
				l++
				r--
			}else if nums[l] + nums[r] >0{
				r--
			}else{
				l++
			}

		}
	}
	return res
}