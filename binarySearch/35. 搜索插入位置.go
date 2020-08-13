package binarySearch

func searchInsert(nums []int, target int) int {
	//找到第一个大于等于目标值的元素
	left,right := 0,len(nums)-1

	//把左指针右移
	for left + 1 <right{
		mid := left + (right- left) / 2
		if nums[mid] <= target{
			left = mid
		}else{
			right = mid
		}
	}

	if nums[left] >= target{
		return left
	} else if nums[right] >= target{
		return right
	}else{
		return right + 1
	}

}
