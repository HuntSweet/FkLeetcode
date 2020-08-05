package binarySearch

func searchRange (A []int, target int) []int {
	//思路：先找到最左边的索引，再找到最右边的索引，然后返回
	left := 0
	right := len(A) - 1
	res := make([]int,2)

	for left + 1 < right{
		mid := left + (right - left) / 2
		if A[mid] >= target{
			right = mid
		}else{
			left = mid
		}
	}

	if A[left] == target{
		res[0] = left
	}else if A[right] == target{
		res[0] = right
	}else{
		res[0],res[1] = -1,-1
		return res
	}

	left = 0
	right = len(A) - 1
	for left + 1 < right{
		mid := left + (right - left) / 2
		if A[mid] > target{
			right = mid
		}else{
			left = mid
		}
	}

	if A[right] == target{
		res[1] = right
	}else if A[left] == target{
		res[1] = left
	}else{
		res[0],res[1] = -1,-1
		return res
	}

	return res
}

