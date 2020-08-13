package main

import (
	"fmt"
	"math"
)

func getCount(n int) (res float64 ){
	//思路：根据n奇偶性判断符号

	for i:=1;i<=2*n;i++{
		if i % 2 == 0{
			res -= float64(1)/(float64(5)*float64(i))
		}else{
			res += float64(1)/(float64(5)*float64(i))
		}
	}


	return CorrectFloat64PrecisionN(res, 4)
}

func CorrectFloat64PrecisionN(f float64, n int) float64 {
	n10 := math.Pow10(n)
	inst := math.Trunc((f+0.5/n10)*n10) / n10
	return inst
}

func main()  {
	var n int
	fmt.Scan(&n)
	//0.1000
	fmt.Printf("%.4f",getCount(n))
}