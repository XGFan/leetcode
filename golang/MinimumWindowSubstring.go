package main

import "fmt"

func minWindow(s string, t string) string {
	try := [128]int{}
	for _, v := range t {
		try[v]++
	}
	var result string
	for l, r, count := 0, 0, 0; r < len(s); r++ {
		try[s[r]]--
		if try[s[r]] >= 0 {
			count++
		}
		if count < len(t) || try[s[r]] != 0 {
			continue
		}
		//li向右移动，直到不满足条件
		for {
			try[s[l]]++
			if l > r || try[s[l]] > 0 {
				break
			} else {
				l++
			}
		}
		if result == "" || len(result) > r+1-l {
			result = s[l : r+1]
		}
		l++
	}
	return result
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("wgrgserthsbv", "a"))
	fmt.Println(minWindow("aa", "aa"))
	fmt.Println(minWindow("cabefgecdaecf", "cae"))
}
