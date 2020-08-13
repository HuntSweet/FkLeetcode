package binarySearch



//简化的思路，拉成一维数组
func searchMatrix(matrix [][]int, target int) bool {
	// 思路：将2纬数组转为1维数组 进行二分搜索
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	start := 0
	end := row*col - 1
	for start+1 < end {
		mid := start + (end-start)/2
		// 获取2纬数组对应值
		val := matrix[mid/col][mid%col]
		if val > target {
			end = mid
		} else if val < target {
			start = mid
		} else {
			return true
		}
	}
	if matrix[start/col][start%col] == target || matrix[end/col][end%col] == target{
		return true
	}
	return false
}
//这里的思路其实考虑的复杂了，但还是对的

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) <= 0{
		return false
	}

	left := 0
	right := len(matrix) - 1
	ll := len(matrix[0])
	for left <= right{
		mid := left + (right-left)/2
		temp := searchInsert(matrix[mid],target)
		if temp == -1{
			return true
			//注意这里不能使用0 < temp < ll
		}else if 0 < temp && temp < ll{
			return false
		}else if temp == 0{
			right = mid
		}else{
			left = mid
		}
	}

	return false
}

//通过返回要插入的位置得到是否存在，以便进行下一步操作
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

	if nums[left] > target{
		return left
	} else if nums[left] == target{
		return -1
	}else if nums[right] > target{
		return right
	}else if nums[right] == target{
		return -1
	}else{
		return right + 1
	}

}
