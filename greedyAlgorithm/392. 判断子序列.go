package greedyAlgorithm

func isSubsequence(s string, t string) bool {
	//子序列的话，所有字符都是按顺序出现在另一个字符串中
	//使用双指针，一个在s，一个在t,找到第一个后移，都找到的话，返回true
	i,j := 0,0
	for i<len(s) && j < len(t){
		if s[i] == t[j]{
			i++
		}
		j++
	}

	if i == len(s){
		return true
	}
	return false
}
