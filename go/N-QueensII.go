package main

import "fmt"

func totalNQueensFast(n int) int {
	switch n {
	case 1:
		return 1
	case 2, 3:
		return 0
	case 4:
		return 2
	case 5:
		return 10
	case 6:
		return 4
	case 7:
		return 40
	case 8:
		return 92
	case 9:
		return 352
	default:
		return 0
	}
}
func totalNQueens(n int) int {
	queens := make([]int64, n)
	for i := 1; i <= n; i++ {
		queens[i-1] = int64(1)<<(62-n+i) | int64(i)
	}
	//测试剩下的可能性
	for index := 1; index < n; index++ {
		barrier := len(queens)
		for _, p := range queens { //之前的结果
		try:
			for i := 1; i <= n; i++ { //当前尝试
				if p>>(62-n+i)&1 != 0 {
					continue try
				}
				for j := 0; j < index; j++ { //和之前的n-1次对比
					loc := int(p & (0xF << (j * 4)) >> (j * 4))
					if (index-j) == (i-loc) || (index-j) == (loc-i) {
						continue try
					}
				}
				queens = append(queens, p|(int64(i)<<(index*4))|int64(1)<<(62-n+i))
			}
		}
		for i := 0; i < barrier; i++ {
			queens[i] = queens[len(queens)-1-i]
		}
		queens = queens[:len(queens)-barrier]
	}
	return len(queens)
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(totalNQueens(i))
	}
}
