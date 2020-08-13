package main

import (
	"fmt"
	"math"
)

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

func main()  {
	var n int
	fmt.Scan(&n)

	l := make([]int,0)
	for i:=0;i<n;i++{
		var t int
		fmt.Scan(&t)
		l = append(l,t)
	}


}
