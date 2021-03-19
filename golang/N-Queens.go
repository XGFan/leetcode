package main

import "fmt"

func solveNQueens(n int) [][]string {
	queens := make([]int64, n)
	bytes := make([]byte, n)
	st := make([]string, n)
	for i := 1; i <= n; i++ {
		bytes[i-1] = '.' //默认全部是.
		queens[i-1] = int64(1)<<(62-n+i) | int64(i)
	}
	//这个循环是第一次尝试
	for i := 0; i < n; i++ {
		bytes[i] = 'Q'
		st[i] = string(bytes) //这儿是一个模板
		bytes[i] = '.'
	}
	//测试剩下的可能性
	for index := 1; index < n; index++ {
		barrier := len(queens)     //现在的queens都是上一轮的结果
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
				//把这轮的结果加到queens
				queens = append(queens, p|(int64(i)<<(index*4))|int64(1)<<(62-n+i))
			}
		}
		//根据barrier，把更早的去掉
		for i := 0; i < barrier; i++ {
			queens[i] = queens[len(queens)-1-i]
		}
		queens = queens[:len(queens)-barrier]
	}
	result := make([][]string, len(queens))
	for i, v := range queens {
		t := make([]string, n)
		for j := 0; j < n; j++ {
			loc := int(v & (0xF << (j * 4)) >> (j * 4))
			t[j] = st[loc-1]
		}
		result[i] = t
	}
	return result
}

func solveNQueens2(n int) [][]string {
	queens := make([][]uint8, 0, n)
	for i := 1; i <= n; i++ {
		queens = append(queens, []uint8{uint8(i)})
	}
	for currCol := 1; currCol < n; currCol++ { //当前遍历的第X行
		barrier := len(queens)
		for i := 0; i < barrier; i++ {
			prev := queens[i]
		try:
			for currPos := 1; currPos <= n; currPos++ {
				for prevCol := 0; prevCol < currCol; prevCol++ {
					prevPos := prev[prevCol]
					if int(prevPos) == currPos || currCol+currPos == int(prevPos)+prevCol || currCol+int(prevPos) == currPos+prevCol {
						continue try
					}
				}
				ints := make([]uint8, currCol, currCol+1)
				copy(ints, prev)
				ints = append(ints, uint8(currPos))
				queens = append(queens, ints)
			}
		}
		for i := 0; i < len(queens)-barrier; i++ {
			queens[i] = queens[barrier+i]
		}
		for i := len(queens) - barrier; i < (len(queens)); i++ {
			queens[i] = nil
		}
		queens = queens[:len(queens)-barrier]
	}
	bytes := make([]byte, n, n)
	st := make([]string, 0, n)
	for i := 0; i < n; i++ {
		bytes[i] = '.'
	}
	for i := 0; i < n; i++ {
		bytes[i] = 'Q'
		st = append(st, string(bytes))
		bytes[i] = '.'
	}
	result := make([][]string, 0, n)
	for _, v := range queens {
		t := make([]string, 0, n)
		for i := 0; i < n; i++ {
			t = append(t, st[v[i]-1])
		}
		result = append(result, t)
	}
	return result
}

func QueensN2(n int) []int64 {
	queens := make([]int64, 0, n)
	for i := 1; i <= n; i++ {
		queens = append(queens, int64(1)<<(62-n+i)|int64(i))
	}
	for index := 1; index < n; index++ {
		barrier := len(queens)
		for _, p := range queens { //之前的结果
		try:
			for i := 1; i <= n; i++ { //当前尝试
				if int64(1)<<(62-n+i)&p != 0 {
					continue
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
		for i := 0; i < len(queens)-barrier; i++ {
			queens[i] = queens[barrier+i]
		}
		for i := len(queens) - barrier; i < len(queens); i++ {
			queens[i] = 0
		}
		queens = queens[:len(queens)-barrier]
	}
	return queens
}

func QueensN(maybe []int64, index, total int) []int64 {
	if index > total {
		return maybe
	}
	barrier := len(maybe)
	if index == 1 {
		for i := 1; i <= total; i++ {
			maybe = append(maybe, int64(1)<<(62-total+i)|int64(i))
		}
	} else {
		for _, p := range maybe { //之前的结果
		try:
			for i := 1; i <= total; i++ { //当前尝试
				if int64(1)<<(62-total+i)&p != 0 {
					continue
				}
				for j := 0; j < index-1; j++ { //和之前的n-1次对比
					loc := int(p & (0xF << (j * 4)) >> (j * 4))
					if (index-j-1) == (i-loc) || (index-j-1) == (loc-i) {
						continue try
					}
				}
				maybe = append(maybe, p|(int64(i)<<((index-1)*4))|int64(1)<<(62-total+i))
			}
		}
		for i := 0; i < len(maybe)-barrier; i++ {
			maybe[i] = maybe[barrier+i]
		}
		for i := len(maybe) - barrier; i < len(maybe); i++ {
			maybe[i] = 0
		}
	}
	return QueensN(maybe[:len(maybe)-barrier], index+1, total)
}

func Queens(k, n int) []int64 {
	if n == 1 {
		t := make([]int64, 0, k)
		for i := 1; i <= k; i++ {
			t = append(t, int64(1)<<(62-k+i)|int64(i))
			//fmt.Printf("%064b\n", int64(1)<<(62-k+i)|int64(i))
		}
		return t
	}
	queens := Queens(k, n-1)
	result := make([]int64, 0, len(queens)/2)
	for _, p := range queens { //之前的结果
	try:
		for i := 1; i <= k; i++ { //当前尝试
			if int64(1)<<(62-k+i)&p != 0 {
				continue
			}
			for j := 0; j < n-1; j++ { //和之前的n-1次对比
				loc := int(p & (0xF << (j * 4)) >> (j * 4))
				if loc == 0 {
					fmt.Printf("%d in %064b\n", j, p)
				}
				if (n-j-1) == (i-loc) || (n-j-1) == (loc-i) {
					continue try
				}
			}
			result = append(result, p|(int64(i)<<((n-1)*4))|int64(1)<<(62-k+i))
			//fmt.Printf("%064b\n", int64(1)<<i<<(63-k)|int64(i))
		}
	}
	//for _, v := range result {
	//	fmt.Printf("%064b\n", v)
	//}
	//fmt.Printf("%d : %d \n", n, len(result))
	return result
}

//9*9
//1248

func main() {
	//p := int64(7)
	//fmt.Println(7)
	//fmt.Println(0xF)
	//fmt.Println(0xF << 0)
	//fmt.Println(p & 0xF >> 0)
	//j := 0
	//loc := (p & (0xF << (j * 4))) >> (j * 4)
	//fmt.Println(loc)
	//111111
	result := solveNQueens(2)
	fmt.Printf("%v\n", result)
	fmt.Printf("%v\n", len(result))
	//fmt.Println(63 & 0xF)
	//fmt.Println(0xF << 4)
	//i := uint32(1 << 31)
	//fmt.Println(int32(-5) &^ i)
	//var i int = -10
	//fmt.Println(i << 1)
	//fmt.Printf("%#b\n", -10)
	//fmt.Printf("%#b\n", 10)
	//fmt.Println(-10 & 10)
	//fmt.Println(10 & -10)
}
