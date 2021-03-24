package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	if len(word1)*len(word2) == 0 {
		return len(word1) + len(word2)
	}
	cache := make([]int, len(word2)+1)
	t := 0
	for x := 0; x <= len(word1); x++ {
		for y := range cache {
			if x == 0 {
				cache[y], t = y, cache[y]
			} else if y == 0 {
				cache[y], t = x, cache[y]
			} else {
				if word1[x-1] == word2[y-1] {
					cache[y], t = min(cache[y-1]+1, cache[y]+1, t), cache[y]
				} else {
					cache[y], t = min(cache[y-1]+1, cache[y]+1, t+1), cache[y]
				}
			}
		}
	}
	return cache[len(word2)]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

func main() {
	fmt.Println(minDistance("intention", "execution") - 5)
	fmt.Println(minDistance("horse", "ros") - 3)
	fmt.Println(minDistance("a", "a") - 0)
}
