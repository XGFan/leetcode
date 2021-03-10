package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	f, t := newInterval[0], newInterval[1]
	for j := 0; j < len(intervals); j++ {
		ints := intervals[j]
		if t < ints[0] {
			//...f...t...o...o...
			//insert and continue
			intervals = append(intervals, newInterval) //增加空间
			copy(intervals[j+1:], intervals[j:])       //挪动
			intervals[j] = newInterval                 //赋值
			return intervals
		} else if t < ints[1] {
			if f < ints[0] {
				//...f...o...t...o...
				ints[0] = f
				return intervals
			} else {
				//...o...f...t...o...
				return intervals
			}
		} else {
			if f <= ints[0] {
				//...f...o...o...t...
				ints[0] = f
				ints[1] = t
				copy(intervals[j:], intervals[j+1:])
				return insert(intervals[:len(intervals)-1], ints)
			} else if f <= ints[1] {
				//...o...f...o...t...
				ints[1] = t
				copy(intervals[j:], intervals[j+1:])
				return insert(intervals[:len(intervals)-1], ints)
			} else {
				//...o...o...x...x...
				//下一轮处理
			}
		}
	}
	return append(intervals, newInterval)
}

func main() {
	fmt.Println(insert([][]int{}, []int{5, 7}), "=", [][]int{{5, 7}})
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 3}), "=", [][]int{{1, 5}})
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 7}), "=", [][]int{{1, 7}})

}
