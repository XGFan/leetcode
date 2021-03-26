package main

import "fmt"

func minWindow(s string, t string) string {
	try := [128]int{}
	for _, v := range t {
		try[v]++
	}
	var indexes []int
	for i := range s {
		if try[s[i]] != 0 {
			indexes = append(indexes, i)
		}
	}
	var li, ri int
	var result string
loop:
	for ri < len(indexes) && li < len(indexes) {
		try[s[indexes[ri]]]--
		for _, v := range try {
			if v > 0 {
				ri++
				continue loop
			}
		}
		if result == "" || len(result) > indexes[ri]+1-indexes[li] {
			result = s[indexes[li] : indexes[ri]+1]
		}
		//fmt.Println(s[l : r+1])
		try[s[indexes[li]]]++
		try[s[indexes[ri]]]++
		li++
	}
	return result
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("wgrgserthsbv", "a"))
	fmt.Println(minWindow("aa", "aa"))
}
