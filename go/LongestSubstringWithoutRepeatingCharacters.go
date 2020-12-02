package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	latestRepeatFrom := 0
	maxLen := 0
	for i, b := range s {
		lastSeen := m[b]
		if lastSeen > latestRepeatFrom {
			//该字符上一次出现，在最近一次重复（from...to）的from之后
			//这样计算长度，就可以从lastSeen-now，只会包含一个to，不会包含from
			latestRepeatFrom = lastSeen
		}
		//该字符上一次出现，在最近一次重复（from...to）的from之前
		//那么只能从from...now，否则会包含from和to
		length := i - latestRepeatFrom + 1
		//找到最长的不重复字符串
		if length > maxLen {
			maxLen = length
		}
		m[b] = i + 1
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring(" "))
}

// f1 f2 t1 t2
// max(f2-0,t1-f1,t2-f2,len-t1)

// f1 t1 f2 t2
// max(t1-0,f2-f1,t2-t1,len-f2)

// f1 f2 t2 t1
// max(f2-0,t2-f1,t1-f2,len-t2)
