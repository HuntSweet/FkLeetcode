package main

import "fmt"
//思路：从最后一个开始，看看构成回文的最大长度，然后后面加上除去这个最大长度的反转的字符串
func isHuiwen(s string) bool {
	n := len(s)

	for i:=0;i<n/2;i++{
		if s[i] != s[n-1-i]{
			return false
		}
	}

	return true

}
func dealhw(s string) string {
	n := len(s)
	//记录下最长的索引位置
	res := n
	t := ""
	for i:=n-1;i>=0;i--{

		if isHuiwen(s[i:]){
			if res > i{
				res = i
			}
		}
	}
	//fmt.Println(res)
	for j :=len(s[:res])-1;j>=0;j--{
		t += string(s[j])
	}
	return s + t
}

func main()  {
	var s string
	fmt.Scan(&s)
	if isHuiwen(s){
		fmt.Println(s)
	}else{
		fmt.Println(dealhw(s))
	}

}