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
	var result string
loop:
	for li, ri := 0, 0; ri < len(indexes) && li < len(indexes); ri++ {
		try[s[indexes[ri]]]--
		if li == 0 {
			for _, v := range try {
				if v > 0 {
					continue loop
				}
			}
			//fmt.Println("first", indexes[li], indexes[ri], try[s[indexes[ri]]])
		} else {
			if try[s[indexes[ri]]] != 0 {
				continue
			}
			//fmt.Println("other", indexes[li], indexes[ri], try[s[indexes[ri]]])
		}
		//li向左移动，直到不满足条件
		for {
			try[s[indexes[li]]]++
			if li > ri || try[s[indexes[li]]] > 0 {
				break
			} else {
				li++
			}
		}
		//fmt.Println("finally", indexes[li], indexes[ri], try[s[indexes[ri]]])
		if result == "" || len(result) > indexes[ri]+1-indexes[li] {
			result = s[indexes[li] : indexes[ri]+1]
		}
		li++
		//fmt.Println(result)
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
