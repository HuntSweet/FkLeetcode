package main

func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0{
		return 0
	}
	clos := len(grid[0])

	var nums int
	for i:=0;i<rows;i++{
		for j:=0;j<clos;j++{
			if grid[i][j] == '1'{
				nums++
				dfs(&grid,i,j)
			}
		}
	}

	return nums
}

func dfs(grid *[][]byte,r,c int){
	temp := *grid
	rows := len(*grid)
	clos := len(temp[0])

	temp[r][c] = '0'
	if r - 1 >=0 && temp[r-1][c] == '1'{
		dfs(grid,r-1,c)
	}
	if r+1 < rows && temp[r+1][c] == '1'{
		dfs(grid,r+1,c)
	}

	if c - 1 >=0 && temp[r][c-1] == '1'{
		dfs(grid,r,c-1)
	}

	if c + 1 < clos && temp[r][c+1] == '1'{
		dfs(grid,r,c+1)
	}
}