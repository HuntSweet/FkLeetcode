package main

import (
	"fmt"
	"strconv"
)

func getNums(n int) (res [][]int) {

	for i:=1;i<=n;i++{

		if re,ok := panduan(n,i);ok{
			temp := []int{0,0}
			temp[0] = re
			temp[1] = re*4
			res = append(res,temp)
		}

	}

	return res
}

func panduan(n,m int) (int,bool) {
	t :=  m * 4
	if n < t{
		return 0,false
	}
	ts := strconv.Itoa(t)
	ls := ""
	for i:=len(ts)-1;i>=0;i--{
		ls += string(ts[i])
	}
	ms := strconv.Itoa(m)
	if ms == ls{
		return m,true
	}

	return 0,false
}
func main()  {
	var n int



	fmt.Scan(&n)

	//if n >=2178 && n<21978{
	//	fmt.Println(1)
	//	fmt.Println(2178,8712)
	//}else if n < 2178{
	//	fmt.Println(0)
	//}else if n>=21978 && n<219978{
	//	fmt.Println(2)
	//	fmt.Println(2178,8712)
	//	fmt.Println(21978,87912)
	//}else if n>=219978 && n<2199978{
	//	fmt.Println(3)
	//	fmt.Println(2178,8712)
	//	fmt.Println(21978,87912)
	//	fmt.Println(219978,879912)
	//}else if n >= 2199978 && n < 21782178{
	//	fmt.Println(4)
	//	fmt.Println(2178,8712)
	//	fmt.Println(21978,87912)
	//	fmt.Println(219978,879912)
	//	fmt.Println(2199978,8799912)
	//}else if n >= 21782178 && n < 21999978{
	//	fmt.Println(5)
	//	fmt.Println(2178,8712)
	//	fmt.Println(21978,87912)
	//	fmt.Println(219978,879912)
	//	fmt.Println(2199978,8799912)
	//	fmt.Println(21782178,87128712)
	//}else{
	//	fmt.Println(0)
	//}

	res := getNums(n)
	fmt.Println(len(res))
	for _,v := range getNums(n){
		fmt.Println(v[0],v[1])
	}

}
