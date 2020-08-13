package main

import "fmt"

func letterCombinations(digits string) []string {
	//思路：采用递归，遍历每个数字字符，找到每个数字字符对应的字母串，相加
	//结束条件：长度等于digits字符串长度
	res := make([]string,0)
	lm := []string{" ","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
	var f func(digits string,index int,s string)
	f = func (digits string,index int,s string){
		if len(s) == len(digits){
			res = append(res,s)
			return
		}

		char := digits[index]
		nums := char - '0'

		for i:=0;i<len(lm[nums]);i++{

			f(digits,index+1,s+string(lm[nums][i]))
		}
		return
	}
	s := ""
	f(digits,0,s)
	return res
}

func main()  {
	fmt.Println(letterCombinations("23"))
}

