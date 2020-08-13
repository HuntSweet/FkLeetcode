//假设小杰头上还有N个bug没修，他每喝一杯咖啡（喝咖啡的时间忽略不计），就能让自己一小时内的debug效率提升到原来的A倍，一小时内重复喝没用，
//最多只能喝X杯，喝太多了晚上会睡不着，并且为了保证可持续发展，每天最多只能专注工作8个小时，而在没喝咖啡的状态下解决每个bug所需要的时间为
//t1,t2...tN 分钟。
//
//现在是早上8点，小杰马上要开始一天的工作了，他计划按从1-N的顺序修复他头上这些Bug，你能帮他计算出他今天能在8小时内解完所有bug吗？
//如果能，最少需要多长时间？
//
//
//
//输入描述
//先按顺序输入三个正整数 N A, X，其中N表示bug的总数，A表示喝一杯咖啡在一小时内debug效率提升的倍数，X表示最多可以喝的咖啡数目。
//（1 <= N <= 100, 1 <= A <= 8, 1 <= X <= 8）
//
//再按顺序输入N个正整数，第i个正整数ti表示解决第i个bug需要的分钟数，第N个正整数tN表示解决第N个bug需要的分钟数，
//（1<=i<=N, 1 <= ti <= 1000）
//
//输出描述
//输出一个数字T，如果不能解完所有bug，则输出0，如果可以，则输出最少需要的分钟数（T为正整数，如不满一分钟则按一分钟计算，
//一旦超过8小时则认为不能解完，所以T最大为480）
package main

import (
	"fmt"
	"io"
	"sort"
)

type TrieNode struct {
	word     string
	children [26]*TrieNode
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	for _, w := range words {
		node := root
		for _, c := range w {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &TrieNode{}
			}
			node = node.children[c-'a']
		}
		node.word = w
	}

	result := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for k,_ := range board[i] {
			dfs(i, k, board, root, &result)
		}
	}
	return result
}

func dfs(i, j int, board [][]byte, node *TrieNode, result *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[i]) {
		return
	}
	c := board[i][j]
	if c == '#' || node.children[c-'a'] == nil {  //访问过了或者字典中没有
		return
	}
	node = node.children[c-'a']

	if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""  // 防止重复添加
	}

	board[i][j] = '#'
	dfs(i+1, j, board, node, result)
	dfs(i-1, j, board, node, result)
	dfs(i, j+1, board, node, result)
	dfs(i, j-1, board, node, result)
	board[i][j] = c
}
//4
//atmb
//fueg
//lyin
//pkjd
//4
//dji
//drone
//flying
//ideas
func main()  {
	//创建二维数组
	board := [][]byte{}
	words := []string{}
	var m,n int
	var err error
	for {
		//第一行要输入的
		_,err = fmt.Scan(&m)
		if err == io.EOF{
			break
		}


		//后面每行要输入的
		for i := 0; i < m; i++ {
			var t string
			fmt.Scan(&t)
			l := []byte(t)
			board = append(board,l)
		}

		fmt.Scan(&n)
		for i := 0; i < n; i++ {
			var t string
			fmt.Scan(&t)

			words = append(words,t)
		}

		if m <=0 || n <=0{
			fmt.Println("")
		}
		//fmt.Println(board,words)
		res := findWords(board,words)
		sort.Strings(res)
		for _,v := range res{
			fmt.Println(v)
		}
	}

}