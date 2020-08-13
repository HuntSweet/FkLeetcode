package main

import (
	"fmt"
	"io"
)




func main()  {
	var m,n int
	var err error
	for {
		//第一行要输入的
		_,err = fmt.Scan(&m,&n)
		if err == io.EOF{
			break
		}
		fmt.Println(m,n)

		//后面每行要输入的
		for i := 0; i < n; i++ {
			var a,b,c int
			fmt.Scan(&a)
			fmt.Scan(&b)
			fmt.Scan(&c)
			fmt.Println(a,b,c)
		}

	}

}
