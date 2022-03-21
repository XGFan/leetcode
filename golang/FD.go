package main

import "fmt"

func main() {
	//fmt.Println(700000 / 360)
	//fmt.Println(700000 * 3.25 / 100 / 12)
	//fmt.Println(700000/360 + 70*10000*3.25/100/12)
	//每月公积金4000
	pay()
}

func pay() {
	total := 700000.0
	rate := 3.25
	account := 200000.0
	income := 9000.0
	var i int
	f := total / 360
	for i = 1; account <= total; i++ {
		shouldPay := f + (total * rate / 100 / 12)
		total -= total / 360
		account += income - shouldPay
		fmt.Println(i, "total:", total, "should", shouldPay, "account:", account)
		if i%12 == 0 {
			total -= account
			account = 0
			fmt.Println("Reload", "total:", total, "account:", account)
		}
	}
	fmt.Println(i)
}
