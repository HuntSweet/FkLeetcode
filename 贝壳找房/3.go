package main

import "fmt"
//思路：找出所有连续区间并计算最大值，同时保存最大值对应的子区间长度
func getSub(l []int) (res int) {

	temp := 0
	res = 1<<32
	n := len(l) -1
	var f func(index int,re []int)
	//index表示不考虑前面已经选取的数，避免重复
	f = func(index int,re []int){
		if len(re) == 2{
			ll := append([]int{},re...)
			//fmt.Println(ll)
			pp := geth(ll[0],ll[1],l)
			if pp>temp{
				res = min(res,ll[1]-ll[0]+1)
				temp = pp
			}

			return
		}

		//[i,n]分别区间两端分别剪枝
		for i:=index;i<=n;i++{

			f(i+1,append(re,i))
		}
	}
	f(0,[]int{})
	for i:=0;i<n;i++{
		pp := geth(i,i,l)
		if pp>temp{
			res = min(res,1)
			temp = pp
		}
	}

	return res
}

func min(x,y int)int  {
	if x > y{
		return y
	}
	return x
}

func geth(l,r int,ls []int) (re int) {
	re = ls[l]
	for l+1<=r{

		re |= ls[l+1]
		fmt.Println(re)
		l++
	}
	return
}
func main()  {
	//var n int
	//fmt.Scan(&n)
	//l := make([]int,0)
	//for i:=0;i<n;i++{
	//	var t int
	//	fmt.Scan(&t)
	//	l = append(l,t)
	//}
	l := []int{1,2,3}
	fmt.Println(getSub(l))

}
