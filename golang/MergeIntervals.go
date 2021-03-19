package main

import "fmt"

func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
l:
	for i := range intervals {
		v := intervals[i]
		f := v[0]
		t := v[1]
		for j := 0; j < len(result); j++ {
			ints := result[j]
			if t < ints[0] {
				//...f...t...o...o...
				//insert and continue
				c := make([]int, 2)
				copy(c, intervals[i])
				result = append(result, []int{}) //增加空间
				copy(result[j+1:], result[j:])   //挪动
				result[j] = c                    //赋值
				continue l
			} else if t < ints[1] {
				if f < ints[0] {
					ints[0] = f
					continue l
					//...f...o...t...o...
				} else {
					continue l
					//...o...f...t...o...
				}
			} else {
				if f <= ints[0] {
					//...f...o...o...t...
					ints[0] = f
					ints[1] = t
					continue l
				} else if f <= ints[1] {
					//...o...f...o...t...
					ints[1] = t
					continue l
				} else {
					//...o...o...x...x...
					//下一轮处理
				}
			}
		}
		c := make([]int, 2)
		copy(c, v)
		result = append(result, c)
	}
	if len(result) != len(intervals) {
		return merge(result)
	} else {
		return result
	}
}

func main() {
	//fmt.Println(merge([][]int{{1, 4}, {4, 5}}), [][]int{{1, 5}})
	//fmt.Println(merge([][]int{{4, 5}, {0, 0}}), [][]int{{0, 0}, {4, 5}})
	//fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}), [][]int{{1, 6}, {8, 10}, {15, 18}})
	fmt.Println(merge([][]int{{4, 5}, {2, 4}, {4, 6}, {3, 4}, {0, 0}, {1, 1}, {3, 5}, {2, 2}}), [][]int{{1, 6}, {8, 10}, {15, 18}})
}
