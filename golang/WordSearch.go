package main

import "fmt"

func exist(board [][]byte, word string) bool {
	t := false
	for i := range board {
		for j := range board[i] {
			t = t || check(0, board, i, j, word)
		}
	}
	return t
}

func check(t int, board [][]byte, i, j int, word string) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}
	if word[0] != board[i][j] {
		return false
	}
	if t&(1<<(i+j*len(board))) != 0 {
		return false
	}
	t |= 1 << (i + j*len(board))
	return check(t, board, i-1, j, word[1:]) ||
		check(t, board, i+1, j, word[1:]) ||
		check(t, board, i, j-1, word[1:]) ||
		check(t, board, i, j+1, word[1:])
}

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
}
