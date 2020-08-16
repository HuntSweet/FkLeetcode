package main

import (
	"fmt"
	"io"
	//"strings"
)

func getNum(ls []string) int {
	//找到第一个，然后找到结尾+1，再从新计算+1
	count := 0
	l,r := 0,0

	for r<len(ls){

		if ls[r] == ls[l] && r != l{
			count++
			l = r+1
		}
		r++

	}
	return count
}
func main()  {
	var n int
	var err error
	var s string
	ls := make([]string,0)
	for {
		//第一行要输入的
		_,err = fmt.Scan(&n)
		if err == io.EOF{
			break
		}

		//后面每行要输入的
		for i := 0; i < 2*n; i++ {
			fmt.Scan(&s)
			ls = append(ls,s)
		}

		fmt.Println(getNum(ls))
	}

}

