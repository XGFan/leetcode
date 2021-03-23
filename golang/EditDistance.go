package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	if len(word1)*len(word2) == 0 {
		return len(word1) + len(word2)
	}
	cache := make([][]int, len(word1)+1)
	for i := range cache {
		cache[i] = make([]int, len(word2)+1)
	}
	for x := range cache {
		for y := range cache[x] {
			if x == 0 {
				cache[x][y] = y
			} else if y == 0 {
				cache[x][y] = x
			} else {
				if word1[x-1] == word2[y-1] {
					cache[x][y] = min(cache[x][y-1]+1, cache[x-1][y]+1, cache[x-1][y-1])
				} else {
					cache[x][y] = min(cache[x][y-1]+1, cache[x-1][y]+1, cache[x-1][y-1]+1)
				}
			}
		}
	}
	return cache[len(word1)][len(word2)]
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
	fmt.Println(minDistance("intention", "execution"))
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("a", "a"))
}

//intention
//execution
