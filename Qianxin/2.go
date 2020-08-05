package main

/**
 * 返回亲7数个数
 * @param digit int整型一维数组 组成亲7数的数字数组
 * @return int整型
 */
import "math"
var res int
var l int
func genNum(nums []int) bool{
	temp := 0
	for i:=0;i<len(nums);i++{
		temp += nums[i] * int(math.Pow(float64(10),float64(len(nums)-i)))
	}
	if temp % 7 == 0{
		return true
	}
	return false
}

func backtrack(nums,pathNums []int,used []bool){
	if len(nums) == len(pathNums){
		if genNum(pathNums){
			res++
		}
		return
	}
	for i:=0;i<len(nums);i++{
		if !used[i]{
			used[i] = true
			pathNums = append(pathNums,nums[i])
			backtrack(nums,pathNums,used)
			pathNums = pathNums[:len(pathNums)-1]
			used[i] = false
		}

	}
}
func reletive_7( digit []int ) int {
	// write code here
	//回溯法找到全排列的数组，然后将数组中的数字拼接（根据位数），对每个数字进行整除，得到个数
	l = len(digit)
	if l <= 0 {
		return 0
	}
	pathNums := make([]int,0)
	used := make([]bool,l)
	backtrack(digit,pathNums,used)
	return res


}