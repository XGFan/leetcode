package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return mayEqual(s1, s2)
}

func mayEqual(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			s1 = s1[i:]
			s2 = s2[i:]
			break
		}
	}
	left := make([]int, 26)
	right := make([]int, 26)
	lc, rc := 0, 0
	for i := 1; i < len(s1); i++ {
		v1 := s1[i-1] - 'a'
		v21 := s2[i-1] - 'a'
		v22 := s2[len(s2)-i] - 'a'
		if v1 != v21 {
			left[v1]++
			left[v21]--
			if left[v1] == 0 {
				lc--
			} else if left[v1] == 1 {
				lc++
			}
			if left[v21] == 0 {
				lc--
			} else if left[v21] == -1 {
				lc++
			}
		}
		if lc == 0 {
			if mayEqual(s1[0:i], s2[0:i]) && mayEqual(s1[i:], s2[i:]) {
				return true
			}
		}
		if v1 != v22 {
			right[v1]++
			right[v22]--
			if right[v1] == 0 {
				rc--
			} else if right[v1] == 1 {
				rc++
			}
			if right[v22] == 0 {
				rc--
			} else if right[v22] == -1 {
				rc++
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
