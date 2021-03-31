package main

import "fmt"

func numDecodings(s string) int {
	cache := make([]int, 2)
	if s[len(s)-1] != '0' {
		cache[0], cache[1] = 1, 1
	} else {
		cache[0], cache[1] = 0, 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		switch s[i] {
		case '0':
			cache[0], cache[1] = 0, cache[0]
		case '1':
			if i < len(s)-1 {
				cache[0], cache[1] = cache[0]+cache[1], cache[0]
			} else {
				cache[1] = cache[0]
			}
		case '2':
			if i < len(s)-1 && s[i+1] <= '6' {
				cache[0], cache[1] = cache[0]+cache[1], cache[0]
			} else {
				cache[1] = cache[0]
			}
		default:
			cache[1] = cache[0]
		}
	}
	return cache[0]
}

func main() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("1111"))
}
