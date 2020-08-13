package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getDis(x,y []int) int {
	res := 0
	for i:=0;i<len(x);i++{
		if x[i] > y[i]{
			res += x[i] - y[i]
		}else{
			res += y[i] - x[i]
		}

	}

	return res
}

func main()  {
	var n int
	fmt.Scan(&n)

	l1 := make([]int,0)
	l2 := make([]int,0)
	inputR := bufio.NewReader(os.Stdin)
	s,_ := inputR.ReadString('\n')

	ss := strings.Split(s," ")

	for _,v := range ss{
		var tmp int
		if strings.Contains(v,string('\n')){
			tmp,_ = strconv.Atoi(v[:1])
		}else{
			tmp,_ = strconv.Atoi(v)
		}
		l1 = append(l1,tmp)
	}

	inputL := bufio.NewReader(os.Stdin)
	st,_ := inputL.ReadString('\n')
	sss := strings.Split(st," ")

	for _,v := range sss{
		var tmp int
		if strings.Contains(v,string('\n')){
			tmp,_ = strconv.Atoi(v[:1])
		}else{
			tmp,_ = strconv.Atoi(v)
		}

		l2 = append(l2,tmp)
	}
	//fmt.Println(l1)
	//fmt.Println(l2)
	fmt.Println(getDis(l1,l2))
}