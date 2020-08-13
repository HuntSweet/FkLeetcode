package main

import (
	"fmt"
	"io"
)

func main()  {
	var m,n int
	var err error
	fmt.Scan(&m)
	for {
		//第一行要输入的
		_,err = fmt.Scan(&n)
		if err == io.EOF{
			break
		}
		l := make([]int,0)
		//后面每行要输入的
		count := 0
		for i := 0; i < n; i++ {
			var t int
			fmt.Scan(&t)
			l = append(l,t)
			count+=t
		}

		//进行逻辑处理
		fmt.Println(count)
	}
}