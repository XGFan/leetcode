package main

import "fmt"

func addBinary(a string, b string) string {
	a, b = exchange(a, b)
	as := []byte(a)
	bs := []byte(b)
	t := 0
	ft := len(b) - len(a)
	for i := len(a) - 1; i >= 0; i-- {
		bi := ft + i
		if as[i] == '1' && bs[bi] == '1' {
			bs[bi] = byte(48 + t)
			t = 1
		} else if as[i] == '0' && bs[bi] == '0' {
			bs[bi] = byte(48 + t)
			t = 0
		} else {
			bs[bi] = byte(48 + t ^ 1)
		}
	}
	if t == 0 {
		return string(bs)
	} else {
		for i := ft - 1; i >= 0; i-- {
			if bs[i] == '1' {
				bs[i] = '0'
			} else {
				bs[i] = '1'
				return string(bs)
			}
		}
	}
	return string(append([]byte{'1'}, bs...))
}

func exchange(a, b string) (string, string) {
	if len(a) < len(b) {
		return a, b
	} else {
		return b, a
	}
}

func main() {
	fmt.Println(addBinary("1", "11"))
	fmt.Println(addBinary("1", "0"))
	fmt.Println(addBinary("1", "111110"))
	fmt.Println(addBinary("1010", "1011"))
}
