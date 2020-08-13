package main

import (
	"fmt"
	"math"
)
//给定一个正整数N，寻找[2, N]区间内没有素数的最长子区间，如果最长子区间不止一个，输出任意一个子区间即可
//
//
//
//输入描述
//输入一行，包含一个整数，代表待测区间[2,N]的上界N
//
//输出描述
//输出一行，包含两个整数，代表区间[2,N]内没有素数的最长子区间，子区间为闭区间


//判断一个数是不是素数
func issushu(n int)bool {
	if n==2||n==3{
		return true
	}
	if n%6 != 1 && n%6!=5{
		return false
	}
	for i:=5;i<=int(math.Sqrt(float64(n+1)));i+=6{
		if n % i == 0 || n % (i + 2) == 0 {
			return false
		}
	}
	return true
}

func getQujian(n int) (res [2]int){
	//思路：两个指针
	count := 0
	l,r := 2,2
	t := 0
	is := false
	for i:=2;i<=n;i++{
		temp := issushu(i)
		if temp{
			if t%2 == 0{
				l = i
			}else{
				r = i
			}
			t++
		}

		if l > r{
			is,count = max(count,l-r)
			if is{
				res[0] = r
				res[1] = l
			}
		}else{
			is,count = max(count,r-l)
			if is{
				res[0] = l
				res[1] = r
			}
		}

	}

	return
}

func max(x,y int) (bool,int) {
	if x > y{
		return false,x
	}

	return true,y
}

func main()  {
	var n int
	fmt.Scan(&n)

	res := getQujian(n)

	fmt.Println(res[0]+1,res[1]-1)
}