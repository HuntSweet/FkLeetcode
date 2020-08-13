package binarySearch

//模板1
func search1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right{
		mid := left + (right - left)/2
		if nums[mid] > target{
			right = mid - 1
		}else if nums[mid] == target{
			return mid
		}else{
			left = mid + 1
		}
	}

	return -1
}

//模板3
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left+1<right{
		mid := left + (right - left)/2
		if nums[mid] > target{
			right = mid
		}else if nums[mid] == target{
			return mid
		}else{
			left = mid
		}
	}

	if nums[left] == target{
		return left
	}
	if nums[right] == target{
		return right
	}

	return -1
}
