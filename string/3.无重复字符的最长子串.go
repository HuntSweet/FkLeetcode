package string

func lengthOfLongestSubstring(s string) int {
	//思路，设置两个指针组成滑动窗口，并设置一个map存储出现过的字符串，如果字符串的某个值出现过，就得到从该字符处最大的子串长度
	//即两个指针中间的长度
	m := make(map[byte]int)
	res := 0
	pre := -1
	for i:=0;i<len(s);i++{
		//每次都要删除之前的，意味着从这个点重新开始
		if i != 0{
			delete(m,s[i-1])
		}

		for pre + 1 < len(s) && m[s[pre+1]] == 0{
			m[s[pre+1]]++
			pre++
		}
		res = max(res,pre-i+1)
	}

	return res
}

func max(x,y int) int  {
	if x > y{
		return x
	}

	return y
}