package main

import (
	"fmt"
	"io"
)


func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func Paper() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		} else {
			//流式处理，O(1)的空间复杂度
			sum := 0 // 所有试卷的总和
			maxAmount := 0 // 试卷最多的组的试卷数
			temp := 0 // 临时变量，用于接收输入
			for i := 0; i < n; i++ {
				fmt.Scan(&temp)
				maxAmount = max(maxAmount, temp) // 记住试卷最多的组的试卷数
				sum += temp // 将所有试卷加起来，得到总试卷数
			}

			if maxAmount > sum-maxAmount {
				fmt.Println("No")
			} else {
				fmt.Println("Yes")
			}
		}
	}
}

func main()  {
	Paper()
}