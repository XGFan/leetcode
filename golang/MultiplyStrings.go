package main

import "fmt"

func add(bs []byte, i, v int) {
	//fmt.Printf("[%d] %d + %d", i, int(bs[i]) - 48, v)
	if int(bs[i]) != 0 {
		v += int(bs[i]) - 48
	}
	//fmt.Printf(" -> %d\n", v%10)
	bs[i] = byte(v%10) + 48
	if v >= 10 {
		add(bs, i-1, v/10)
	}
}

func multiply(num1 string, num2 string) string {
	bytes := make([]byte, len(num1)+len(num2), len(num1)+len(num2))
	for i := 1; i <= len(num1); i++ {
		for j := 1; j <= len(num2); j++ {
			add(bytes, len(bytes)+1-i-j, (int(num1[len(num1)-i])-48)*(int(num2[len(num2)-j])-48))
		}
	}
	p := 0
	for p < len(bytes)-1 && (bytes[p] == 0 || bytes[p] == 48) {
		p++
	}
	return string(bytes[p:])
}

func main() {
	//fmt.Println(int('0'))
	//fmt.Println(string([]byte{byte(48), byte(49)}))
	//ints := make([]int, 0, 10)
	//fmt.Println(cap(ints))
	//fmt.Println(len(ints))
	//fmt.Println(multiply("999", "999"))
	fmt.Println(multiply("123", "456"))
	//fmt.Println(multiply("3", "3"))
	//fmt.Println(multiply("9", "99"))
	//fmt.Println(multiply("3", "0"))

}
