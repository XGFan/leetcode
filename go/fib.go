package main

import (
	"fmt"
	"math"
	"math/rand"
)

func fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func fibT(n int) int {
	a, b, p, q := 0, 1, 0, 1
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

func expmod(base, exp, m int) int {
	if exp == 0 {
		return 1
	}
	if exp%2 == 0 {
		return int(math.Pow(float64(expmod(base, exp/2, m)), 2)) % m
	} else {
		return (base * expmod(base, exp-1, m)) % m
	}
}

func fastPrime(n, times int) bool {
	if times == 0 {
		return true
	}
	a := rand.Intn(n-1) + 1
	if expmod(a, n, n) == a%n {
		return fastPrime(n, times-1)
	} else {
		return false
	}
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(expmod(67, 8911, 8911))
	//fmt.Println(expmod(577, 997633, 997633))
	//fmt.Println(fastPrime(997633, 10000))
	//fmt.Println(fastPrime(4, 100000))
	//fmt.Println(isPrime(4))
	//fmt.Println(expmod(2, 3, 3))
	//fastPrime(3, 10000)
	//for i := 3; i <= 10000; i++ {
	//	if isPrime(i) != fastPrime(i, 10000) {
	//		fmt.Printf("Wrong :%d  -> %v %v\n", i, isPrime(i), fastPrime(i, 10000))
	//	}
	//}
	sum := 0
	for i := 2; i <= 100; i++ {
		if isPrime(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println(sum)
	//start := time.Now()
	//for i := 0; i < 10_000; i++ {
	//	for j := 1; j <= 1000; j++ {
	//		if fib(j) != fibT(j) {
	//			log.Panic(j)
	//		}
	//	}
	//}
	//fmt.Println(time.Now().Sub(start).Milliseconds())
	//start = time.Now()
	//for i := 0; i < 10_000; i++ {
	//	for j := 1; j <= 1000; j++ {
	//		fibT(j)
	//	}
	//}
	//fmt.Println(time.Now().Sub(start).Milliseconds())
}
