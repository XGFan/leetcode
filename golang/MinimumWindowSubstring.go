package main

import "fmt"

func minWindow(s string, t string) string {
	try := make(map[byte]int, len(t)/2)
	for _, v := range []byte(t) {
		try[v] += 1
	}
	indexes := make([]int, 0)
	for i := range s {
		if _, ok := try[s[i]]; ok {
			indexes = append(indexes, i)
		}
	}
	li, ri := 0, 0
	from, to := 0, 0
	if len(indexes) == 0 {
		return ""
	}
	l, r := indexes[li], indexes[ri]
x:
	for ri < len(indexes) && li < len(indexes) {
		r = indexes[ri]
		l = indexes[li]
		try[s[r]]--
		for _, v := range try {
			if v > 0 {
				ri++
				continue x
			}
		}
		if to == 0 || to-from > r+1-l {
			from, to = l, r+1
		}
		//fmt.Println(s[l : r+1])
		try[s[l]]++
		try[s[r]]++
		li++
	}
	return s[from:to]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("wgrgserthsbv", "a"))
	fmt.Println(minWindow("aa", "aa"))
}
