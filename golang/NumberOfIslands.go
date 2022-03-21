package main

func dotToMore(grid [][]byte, i, j int) {
	if !(i < len(grid) && i >= 0 && j < len(grid[0])) || j < 0 {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = 'x'
	dotToMore(grid, i-1, j)
	dotToMore(grid, i+1, j)
	dotToMore(grid, i, j-1)
	dotToMore(grid, i, j+1)
}
func numIslands(grid [][]byte) int {
	count := 0
	for i, bytes := range grid {
		for j, b := range bytes {
			if b == '1' {
				count++
				dotToMore(grid, i, j)
			}
		}
	}
	return count
}

func main() {
	numIslands([][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	})
}
