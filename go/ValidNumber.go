package main

import "fmt"

//0为正负号
//1为整数
//2为浮点
//3为e
//4为.
//e之后只能为整数
func isNumber(s string) bool {
	bytes := []byte(s)
	last := -1
	hasE := false
	for i, b := range bytes {
		switch b {
		case 'E', 'e':
			if hasE || i == 0 || i == len(bytes)-1 {
				return false
			} else {
				if last == 4 || last == 0 {
					return false
				}
				hasE = true
				last = 3
			}
		case '+', '-':
			if last != -1 {
				if last != 3 {
					return false
				}
				last = 0
			} else {
				last = 0
			}
		case '.':
			if last != -1 {
				if hasE {
					return false
				}
				if last == 0 { //正负号
					if i == len(bytes)-1 {
						return false
					} else {
						last = 4
						continue
					}
				}
				if last == 1 { //之前是整数
					last = 2
					continue
				}
				if last == 2 || last == 4 { //之前是浮点或者小数点
					return false
				}
			}
			if i == len(bytes)-1 {
				return false
			} else {
				last = 4
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			//number
			if last != -1 {
				if last == 0 {
					last = 1
					continue
				}
				if last == 1 || last == 2 { //整数或者浮点
					continue
				}
				if last == 3 {
					last = 1
					continue
				}
				if last == 4 {
					last = 2
					continue
				}
			} else {
				last = 1
			}
		default:
			return false
		}
	}
	return last != 0
}

func main() {
	fmt.Println(isNumber("1"))
	fmt.Println(isNumber("0089"))
	fmt.Println(isNumber("-0.1"))
	fmt.Println(isNumber("+3.14"))
	fmt.Println(isNumber("4."))
	fmt.Println(isNumber("-.9"))
	fmt.Println(isNumber("2e10"))
	fmt.Println(isNumber("-90E3"))
	fmt.Println(isNumber("3e+7"))
	fmt.Println(isNumber("+6e-1"))
	fmt.Println(isNumber("53.5e93"))
	fmt.Println(isNumber("-123.456e789"))
	fmt.Println(isNumber("45.e3"))
	//----
	fmt.Println(!isNumber("abc"))
	fmt.Println(!isNumber("1a"))
	fmt.Println(!isNumber("1e"))
	fmt.Println(!isNumber("e3"))
	fmt.Println(!isNumber("99e2.5"))
	fmt.Println(!isNumber("--6"))
	fmt.Println(!isNumber("-+6"))
	fmt.Println(!isNumber("95a54e53"))
	fmt.Println(!isNumber(".e1"))
	fmt.Println(!isNumber("4e+"))
	fmt.Println(!isNumber("+."))
	fmt.Println(!isNumber("+E3"))

}
