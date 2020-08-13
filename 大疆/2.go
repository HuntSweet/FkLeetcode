package main

import (
	"fmt"
	"io"
)

func main()  {
	//创建二维数组

	var m,n,o int
	var err error
	l := []int{}
	count := 0
	for {
		//第一行要输入的
		_,err = fmt.Scan(&m)
		if err == io.EOF{
			break
		}

		fmt.Scan(&n)
		fmt.Scan(&o)
		//后面每行要输入的
		for i := 0; i < m; i++ {
			var t int
			fmt.Scan(&t)
			l = append(l,t)
			count += t
		}

		if n*o*60 >= count && n <= 6{
			if count % o != 0{
				fmt.Println(count/o+1)
			}else{
				fmt.Println(count/o)
			}

		}else if n*o*60 <= count && count - n*o*60 + n*60 <= 480{
			fmt.Println(count - n*o*60 + n*60)
		}else{
			fmt.Println(0)
		}
	}

}

