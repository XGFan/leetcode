package main

import "fmt"

func mySqrt(x int) int {
	guess := 1
	for {
		if guess*guess <= x && (guess+1)*(guess+1) > x {
			return guess
		} else {
			guess = (guess + (x / guess)) / 2
		}
	}
}

func main() {
	fmt.Println(mySqrt(10) == 3)
	fmt.Println(mySqrt(9) == 3)
	fmt.Println(mySqrt(8) == 2)
	fmt.Println(mySqrt(4) == 2)
	fmt.Println(mySqrt(2) == 1)
	fmt.Println(mySqrt(0) == 0)
}

//(define (sqrt x)
//(define (good-enough? x guess)
//(< (abs (- (* guess guess) x)) 0.001))
//(define (sqrt-iter x guess)
//(if (good-enough? x guess)
//guess
//(sqrt-iter x (improve x guess))))
//(define (improve x guess)
//(/ (+ guess (/ x guess)) 2))
//(sqrt-iter x 1.0))
