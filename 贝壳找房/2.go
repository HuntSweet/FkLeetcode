package main

import (
	"fmt"
	"io"
)

func getRes(m,n int)(res int){
	if m*n % 2== 0{
		return 2
	}
}

func main()  {
	var t int
	fmt.Scan(&t)

	for {
		var m,n int
		_,err := fmt.Scanf("%d %d",&m,&n)
		if err == io.EOF{
			break
		}
		fmt.Println(getRes(m,n))
	}
}
