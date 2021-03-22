package main

import (
	"fmt"
)

func simplifyPath(path string) string {
	var f = 0 //0是none
	strs := make([]byte, 0)
	ls := make([]int, 0)
	for i := 0; i <= len(path); i++ {
		if i == len(path) || path[i] == '/' {
			switch i - f {
			case 1:
				//just jump
			case 2:
				if path[f+1] == '.' {
					//just jump
				} else {
					strs = append(strs, path[f:i]...)
					ls = append(ls, i-f)
				}
			case 3:
				if path[f+1] == '.' && path[f+2] == '.' {
					//upper
					if len(ls) != 0 {
						strs = strs[0 : len(strs)-ls[len(ls)-1]]
						ls = ls[0 : len(ls)-1]
					}
				} else {
					strs = append(strs, path[f:i]...)
					ls = append(ls, i-f)
				}
			default:
				strs = append(strs, path[f:i]...)
				ls = append(ls, i-f)
			}
			f = i
		}
	}
	if len(strs) == 0 {
		return "/"
	}
	return string(strs)
}

func main() {
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
	fmt.Println(simplifyPath("/..."))
	fmt.Println(simplifyPath("/.../"))
	fmt.Println(simplifyPath("/..hidden"))
}
