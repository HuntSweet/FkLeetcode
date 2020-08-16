package main

func lengthOfLongestSubstring(s string) int {
	win := make(map[byte]int)

	left,right,max := 0,0,0

	var temp byte
	//如果
	for right < len(s){
		temp = s[right]
		win[temp]++
		right++
		//缩小窗口
		for win[temp] > 1{
			t := s[left]
			win[t]--
			left++
		}

		if right - left > max{
			max = right -left
		}

	}


	return max
}
