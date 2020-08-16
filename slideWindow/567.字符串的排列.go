package main

func checkInclusion(s1 string, s2 string) bool {
	//主要是看字符串是否出现并连续，连续并相等的话就返回true，不连续的话就缩紧窗口
	need := make(map[byte]int)
	win := make(map[byte]int)
	for i:=0;i<len(s1);i++{
		need[s1[i]]++
	}
	left,right,count := 0,0,0
	var temp byte
	for right < len(s2){
		temp = s2[right]
		right++
		if need[temp] != 0{
			win[temp]++
			if win[temp] == need[temp]{
				count++
			}
		}

		// 当窗口长度大于字符串长度，缩紧窗口
		for right-left >= len(s1){
			//如果都满足，返回true
			if count == len(need){
				return true
			}
			temp = s2[left]
			if need[temp] != 0{
				if need[temp] == win[temp]{
					count--
				}
				win[temp]--
			}
			left++
		}

	}
	return false
}
