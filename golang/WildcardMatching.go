package main

import "fmt"

func isMatch(s string, p string) bool {
	var ns = []N{
		Char{
			maybe: []int{-1},
			len:   1,
		},
	}
	from := 0
	for i := 0; i < len(p); i++ {
		switch p[i] {
		case '*':
			if _, ok := ns[len(ns)-1].(Char); ok {
				ns = append(ns, Pattern{0})
			}
		case '?':
			switch v := ns[len(ns)-1].(type) {
			case Char:
				v.len++
				ns[len(ns)-1] = v
			case Pattern:
				v.min++
				ns[len(ns)-1] = v
			}
			from += 1
		default:
			switch v := ns[len(ns)-1].(type) {
			case Char:
				index := 0
				for _, loc := range v.maybe {
					if loc+v.len < len(s) && s[loc+v.len] == p[i] {
						v.maybe[index] = loc
						index++
					}
				}
				v.maybe = v.maybe[:index]
				if len(v.maybe) == 0 {
					return false
				}
				v.len += 1
				ns[len(ns)-1] = v
			case Pattern:
				maybe := make([]int, 0)
				for j := from; j < len(s); j++ {
					if s[j] == p[i] {
						maybe = append(maybe, j)
					}
				}
				if len(maybe) == 0 {
					return false
				}
				ns = append(ns, Char{maybe, 1})
			}
			from += 1
		}
	}
	maybe := []int{-1}
	fuzzy := false
	limit := 0
	fmt.Printf("%+v  ", ns)
	fmt.Printf("%v  -> ", maybe)
	fixed := true
	for i := range ns {
		switch v := ns[i].(type) {
		case Char:
			if fuzzy {
				newMaybe := make([]int, 0)
				for _, y := range v.maybe {
					if maybe[0]+limit <= y {
						newMaybe = append(newMaybe, y+v.len)
					}
				}
				maybe = newMaybe
			} else {
				index := 0
				for _, x := range maybe {
					flag := false
					for _, y := range v.maybe {
						if x == y {
							flag = true
							break
						}
					}
					if flag {
						maybe[index] = x + v.len
						index++
					}
				}
				maybe = maybe[:index]
			}
			limit = v.len
			fixed = true
			fuzzy = false
		case Pattern:
			fuzzy = true
			fixed = false
			limit = v.min
			if maybe[0]+limit > len(s) {
				return false
			}
		}
		fmt.Printf("%v  -> ", maybe)
		if len(maybe) == 0 {
			return false
		}
	}
	//fmt.Print("[ ", limit, " ", fixed, " ] ")
	if fixed {
		for _, try := range maybe {
			if try == len(s) {
				return true
			}
		}
		return false
	} else {
		return maybe[0] <= len(s)
	}

}

type N interface {
}
type Char struct {
	maybe []int
	len   int
}
type Pattern struct {
	min int
}

func isMatch2(s string, p string) bool {
	isChar := true
	maybe := []int{-1}
	limit := 1
	for i := range p {
		switch p[i] {
		case '*':
			isChar = false
		case '?':
			limit++
		default:
			if isChar {
				index := 0
				for _, loc := range maybe {
					if loc+limit < len(s) && s[loc+limit] == p[i] {
						maybe[index] = loc
						index++
					}
				}
				maybe = maybe[:index]
				limit++
			} else {
				from := maybe[0] + limit
				maybe = maybe[:0]
				for j := from; j < len(s); j++ {
					if s[j] == p[i] {
						maybe = append(maybe, j)
					}
				}
				limit = 1
			}
			isChar = true
		}
		//fmt.Printf("%v -> ", maybe)
		if len(maybe) == 0 {
			return false
		}
	}
	if isChar {
		for _, try := range maybe {
			if try+limit == len(s) {
				return true
			}
		}
		return false
	} else {
		return maybe[0]+limit <= len(s)
	}
}

func main() {
	fmt.Println(!isMatch2("aa", "a"))
	fmt.Println(isMatch2("aa", "*"))
	fmt.Println(!isMatch2("cb", "?a"))
	fmt.Println(isMatch2("adceb", "*a*b"))
	fmt.Println(!isMatch2("acdcb", "a*c?b"))
	fmt.Println(!isMatch2("mississippi", "m??*ss*?i*pi"))
	fmt.Println(isMatch2("adceb", "*a*b"))
	fmt.Println(!isMatch2("b", "?*?"))
	fmt.Println(isMatch2("adceb", "*a*b"))
	fmt.Println(isMatch2("", "*****"))
	fmt.Println(!isMatch2("ab", "*a"))
	fmt.Println(isMatch2("baba", "*a??"))
	fmt.Println(isMatch2("abcabczzzde", "*abc???de*"))
}

//[{maybe:[-1] len:4} {min:0} {maybe:[5] len:2} {min:1} {maybe:[7 10] len:1} {min:0} {maybe:[9] len:2}]
//         m??           *            ss           *?             i             *           pi
//		   miss          _
//  	    [-1]  ->     [3]  ->      [3]  ->      [7]  ->       [7]  ->       [11]  ->     [11]  ->       [11]
