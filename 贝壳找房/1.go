package main

import (
	"fmt"
	"io"
)

func getNum(s string) (res int) {
	if ishuiwen(s){
		return
	}

	i,j := 0,len(s)-1
	for i< j{
		if s[i] != s[j]{
			res++
		}
		i++
		j--
	}
	return res
}

func ishuiwen(s string)bool{

	l := len(s)
	if l == 1 {
		return true
	}
	cycle := l/2
	for i:=0;i<cycle;i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
func main(){
	var n int
	var s string
	for {
		_,err := fmt.Scan(&n)
		if err == io.EOF{
			break
		}
		fmt.Scan(&s)

		fmt.Println(getNum(s))
	}

}
