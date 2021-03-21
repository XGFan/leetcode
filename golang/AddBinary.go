package main

import "fmt"

func addBinary(a string, b string) string {
	a, b = exchange(a, b)
	as := []byte(a)
	bs := []byte(b)
	t := 0
	bytes := make([]byte, 0)
	//bytes := make([]byte, len(b))
	//copy(bytes, bs)
	i2 := len(b) - len(a)
	for i := len(a) - 1; i >= 0; i-- {
		if as[i] == '1' {
			if bs[i2+i] == '1' {
				if t == 0 {
					bytes = append(bytes, '0')
					t = 1
				} else {
					bytes = append(bytes, '1')
					t = 1
				}
			} else {
				if t == 0 {
					bytes = append(bytes, '1')
					t = 0
				} else {
					bytes = append(bytes, '0')
					t = 1
				}
			}
		} else {
			if bs[i2+i] == '1' {
				if t == 0 {
					bytes = append(bytes, '1')
					t = 0
				} else {
					bytes = append(bytes, '0')
					t = 1
				}
			} else {
				if t == 0 {
					bytes = append(bytes, '0')
					t = 0
				} else {
					bytes = append(bytes, '1')
					t = 0
				}
			}
		}
	}
	if t == 0 {
		for i := 0; i < len(bytes)/2; i++ {
			bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
		}
		return string(append(bs[:i2], bytes...))
	} else {
		for i := i2 - 1; i >= 0; i-- {
			if bs[i] == '1' {
				bytes = append(bytes, '0')
			} else {
				bytes = append(bytes, '1')
				for j := 0; j < len(bytes)/2; j++ {
					bytes[j], bytes[len(bytes)-1-j] = bytes[len(bytes)-1-j], bytes[j]
				}
				return string(append(bs[:i], bytes...))
			}
		}
	}
	bytes = append(bytes, '1')
	for j := 0; j < len(bytes)/2; j++ {
		bytes[j], bytes[len(bytes)-1-j] = bytes[len(bytes)-1-j], bytes[j]
	}
	return string(bytes)
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
