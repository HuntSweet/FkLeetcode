package main

func minWindow(s string, t string) string {
	//思路:建立两个字典，一个用于保存T的所有字符与对应数量，一个用于保存滑动窗口中T出现的字符及其数量
	//如果T中所有字符与数量与滑动窗口都一样，那么可以得到

	//滑动窗口的字符集
	win := make(map[byte]int)
	//需要的字符集
	need := make(map[byte]int)

	for _,v := range t{
		need[byte(v)]++
	}

	left,right,start,end,count,res := 0,0,0,0,0,1<<31-1
	var temp byte
	for right < len(s){
		temp = s[right]
		right++
		if _,ok := need[temp];ok{
			win[temp]++
			if win[temp] == need[temp]{
				count++
			}
		}

		//当所有字符数量都匹配之后，开始缩紧窗口
		for count == len(need){
			if right - left < res{
				start = left
				end = right
				res = end - start
			}

			//缩紧窗口直到不满足条件，即count ！= len(need)
			// 左指针指向字符数量和需要的字符相等时，右移之后count值不匹配则减一
			// 因为win里面的字符数可能比较多，如有10个A，但需要的字符数量可能为3
			// 所以在压死骆驼的最后一根稻草时，count才减一，这时候才会跳出循环
			if need[s[left]] != 0{
				if need[s[left]] == win[s[left]]{
					count--
				}
				win[s[left]]--
			}
			left++

		}
	}
	return s[start:end]
}

func min(x,y int) int {
	if x > y{
		return y
	}
	return x
}