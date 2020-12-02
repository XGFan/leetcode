package main

import "fmt"

func myPowIter(factor, x float64, n int) float64 {
	if n == 1 {
		return factor * x
	} else if n == -1 {
		return factor / x
	}
	if n%2 == 0 {
		return myPowIter(factor, x*x, n/2)
	} else {
		if n > 0 {
			return myPowIter(factor*x, x*x, n/2)
		} else {
			return myPowIter(factor/x, x*x, n/2)
		}
	}
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	return myPowIter(1.0, x, n)
}

func main() {
	fmt.Println(myPow(3.0, 2))
	fmt.Println(myPow(3.0, 3))
	fmt.Println(myPow(3.0, -3))
}
