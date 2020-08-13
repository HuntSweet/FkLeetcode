package main

import (
	"fmt"
	"math"
	"strconv"
)

func getNums(m,n int) int  {
	count := 0
	for i:=m;i<=n;i++{
		if isNNN(i){
			count++
		}
	}
	return count
}
//将所有回文素数列出来
func isNNN(n int) bool {
	res := []int{2, 3, 5, 7, 11, 101, 131, 151, 181, 191, 313, 353, 373, 383, 727, 757, 787, 797, 919, 929, 10301, 10501, 10601, 11311, 11411, 12421, 12721, 12821, 13331, 13831, 13931, 14341, 14741, 15451, 15551, 16061, 16361, 16561, 16661, 17471, 17971, 18181, 18481, 19391, 19891, 19991, 30103, 30203, 30403, 30703, 30803, 31013, 31513, 32323, 32423, 33533, 34543, 34843, 35053, 35153, 35353, 35753, 36263, 36563, 37273, 37573, 38083, 38183, 38783, 39293, 70207, 70507, 70607, 71317, 71917, 72227, 72727, 73037, 73237, 73637, 74047, 74747, 75557, 76367, 76667, 77377, 77477, 77977, 78487, 78787, 78887, 79397, 79697, 79997, 90709, 91019, 93139, 93239, 93739, 94049, 94349, 94649, 94849, 94949, 95959, 96269, 96469, 96769, 97379, 97579, 97879, 98389, 98689}
	m := make(map[int]bool)
	for _,v := range res{
		m[v] = true
	}
	s := strconv.Itoa(n)

	for i:=0;i<len(s);i++{
		temp := s[:i] + s[i+1:]

		t,_ := strconv.Atoi(temp)
		if m[t]{
			return true
		}

	}
	return false

}

//判断一个数是不是素数
func issushu(n int)bool {
	if n==2||n==3{
		return true
	}
	if n%6 != 1 && n%6!=5{
		return false
	}
	for i:=5;i<=int(math.Sqrt(float64(n+1)));i+=6{
		if n % i == 0 || n % (i + 2) == 0 {
			return false
		}
	}
	return true
}
//判断回文
func ishuiwen(n int)bool{
	s := strconv.Itoa(n)
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

func main()  {
	var m,n int

	fmt.Scanf("%d %d",&m,&n)

	fmt.Println(getNums(m,n))
}