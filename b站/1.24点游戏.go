package main

import "fmt"

/**
 *
 * @param arr int整型一维数组
 * @return bool布尔型
 */
func Game24Points( arr []int ) bool {
	// write code here
	res := make([][]int,0)
	used := make(map[int]bool)
	var f func(re []int)
	f = func(re []int) {
		if len(re) == len(arr){

			res = append(res,append([]int{},re...))
			return
		}

		for i:=0;i<len(arr);i++{

			if !used[i] == true{
				used[i] = true
				f(append(re,arr[i]))
				used[i] = false
			}


		}
	}
	re := []int{}
	f(re)
	fmt.Println(res,used)
	for _,v := range res{
		if Get24(v,-1,1){
			return true
		}
	}
	return false
}


//错误解法：只考虑了从左到右连续加减乘除，未考虑乘法与除法的优先级
//例如：4 + 10 / 2 + 18是对的，却返回false，除非做一趟全排列，这样复杂度就提升了，不过只有4个数还好

//思路：遍历4个数，执行3次加减乘除运算，如果能得到24就返回，
func Get24(arr []int,count int,res int)bool{
	if len(arr) == count+1 && res == 24{
		return true
	}else if len(arr) == count+1 && res != 24{
		return false
	}
	if count == -1{
		return Get24(arr,count+1,res*arr[count+1])
	}
	return Get24(arr,count+1,res*arr[count+1])||Get24(arr,count+1,res/arr[count+1])||
		Get24(arr,count+1,res+arr[count+1]) || Get24(arr,count+1,res-arr[count+1])
}

func main()  {
	fmt.Println(Game24Points([]int{4,10,5,18}))
}