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
	t := make([]int, 128)
	for _, v := range s1 {
		t[v]++
	}
	for _, v := range s2 {
		t[v]--
		if t[v] < 0 {
			return false
		}
	}
	for s1[0] == s2[0] {
		s1 = s1[1:]
		s2 = s2[1:]
	}
	for i := 1; i < len(s1); i++ {
		if mayEqual(s1[0:i], s2[0:i]) && mayEqual(s1[i:], s2[i:]) {
			return true
		}
		if mayEqual(s1[0:i], s2[len(s2)-i:]) && mayEqual(s1[i:], s2[0:len(s2)-i]) {
			return true
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
