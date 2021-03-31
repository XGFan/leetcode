package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return mayEqual([]byte(s1), []byte(s2))
}

func mayEqual(s1, s2 []byte) bool {
	for len(s1) != 0 && s1[0] == s2[0] {
		s1 = s1[1:]
		s2 = s2[1:]
	}
	if len(s1) == 0 {
		return true
	}
	left := make([]int, 128)
	right := make([]int, 128)
	lc, rc := 0, 0
	for i := 1; i < len(s1); i++ {
		if left[s1[i-1]] == 0 {
			lc++
		} else if left[s1[i-1]] == -1 {
			lc--
		}
		if right[s1[i-1]] == 0 {
			rc++
		} else if right[s1[i-1]] == -1 {
			rc--
		}
		left[s1[i-1]]++
		right[s1[i-1]]++

		if left[s2[i-1]] == 0 {
			lc++
		} else if left[s2[i-1]] == 1 {
			lc--
		}
		if right[s2[len(s2)-i]] == 0 {
			rc++
		} else if right[s2[len(s2)-i]] == 1 {
			rc--
		}
		left[s2[i-1]]--
		right[s2[len(s2)-i]]--

		if lc == 0 {
			if mayEqual(s1[0:i], s2[0:i]) && mayEqual(s1[i:], s2[i:]) {
				return true
			}
		}
		if rc == 0 {
			if mayEqual(s1[0:i], s2[len(s2)-i:]) && mayEqual(s1[i:], s2[0:len(s2)-i]) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(isScramble("great", "rgeat"))
	fmt.Println(isScramble("abcde", "caebd"))
	fmt.Println(isScramble("abc", "bca"))
	fmt.Println(isScramble("eebaacbcbcadaaedceaaacadccd", "eadcaacabaddaceacbceaabeccd"))
}
