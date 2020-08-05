package string

import "strconv"

func addStrings(num1 string, num2 string) string {
	//思路:从后往前遍历，注意进位，取模等就ok了
	l1 := len(num1)
	l2 := len(num2)
	carry := 0
	s := make([]string,0)
	for l1 != 0 || l2 != 0 || carry != 0{
		if l1 != 0{
			l1--
			carry += int(num2[l1] - '0')

		}
		if l2 != 0{
			l2--
			carry += int(num2[l2] - '0')

		}

		s = append(s,strconv.Itoa(carry%10))
		carry /= 10
	}

	ls := len(s)

	res := ""
	for l:=ls-1;l>=0;l--{
		res += s[l]
	}
	return res}
