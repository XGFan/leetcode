package main

import "fmt"

func climbStairs(n int) int {
	a, b, p, q := 1, 1, 0, 1
	for n != 0 {
		if n%2 != 0 {
			a, b = a*p+b*q, a*q+b*p+b*q
			n--
		} else {
			p, q = p*p+q*q, q*q+(p*q<<1)
			n = n >> 1
		}
	}
	return a
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
}
