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
	lastIndex := len(bytes) - 1
	for i, b := range bytes {
		switch b {
		case 'E', 'e':
			if hasE || i == 0 || i == lastIndex || last == 4 || last == 0 {
				return false
			} else {
				hasE = true
				last = 3
			}
		case '+', '-':
			switch last {
			case -1, 3:
				last = 0
			default:
				return false
			}
		case '.':
			if hasE {
				return false
			}
			switch last {
			case -1, 0:
				if i == lastIndex {
					return false
				} else {
					last = 4
				}
			case 1:
				last = 2
			case 2, 4:
				return false
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			switch last {
			case -1, 0, 3:
				last = 1
			case 4:
				last = 2
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
