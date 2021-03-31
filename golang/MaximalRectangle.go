package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	max := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				im, jm := len(matrix), len(matrix[i])
				for ip := i; ip < len(matrix); ip++ {
					if matrix[ip][j] == '0' {
						im = ip
						break
					}
				}
				for jp := j; jp < len(matrix[i]); jp++ {
					if matrix[i][jp] == '0' {
						jm = jp
						break
					}
				}
				if max < im-i {
					max = im - i
				}
				if max < jm-j {
					max = jm - j
				}
				//x, y := im-i, jm-j
				for ip := i; ip < im; ip++ {
					for jp := j; jp < jm; jp++ {
						if matrix[ip][jp] == '0' {
							jm = jp
							break
						} else {
							if max < (ip-i+1)*(jp-j+1) {
								max = (ip - i + 1) * (jp - j + 1)
							}
						}
					}
				}
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}
