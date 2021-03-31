package main

import "fmt"

func numDecodings(s string) int {
	cache := make([]int, len(s)+1)
	cache[len(s)] = 1
	if s[len(s)-1] != '0' {
		cache[len(s)-1] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		switch s[i] {
		case '0':
			cache[i] = 0
		case '1':
			if i < len(s)-1 {
				cache[i] = cache[i+1] + cache[i+2]
			} else {
				cache[i] = cache[i+1]
			}
		case '2':
			if i < len(s)-1 && s[i+1] <= '6' {
				cache[i] = cache[i+1] + cache[i+2]
			} else {
				cache[i] = cache[i+1]
			}
		default:
			cache[i] = cache[i+1]
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
