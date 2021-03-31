package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 1
	}
	switch s[0] {
	case '0':
		return 0
	case '1':
		if len(s) > 1 {
			return numDecodings(s[1:]) + numDecodings(s[2:])
		} else {
			return numDecodings(s[1:])
		}
	case '2':
		if len(s) > 1 && s[1] <= '6' {
			return numDecodings(s[1:]) + numDecodings(s[2:])
		} else {
			return numDecodings(s[1:])
		}
	default:
		return numDecodings(s[1:])
	}
}

func main() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("1111"))
}
