package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	l := len(intervals)
	fi, ti := l, l
	for j := 0; j < l; j++ {
		if fi == l && ti != l {
			break
		}
		ints := intervals[j]
		if fi == l {
			if newInterval[0] <= ints[0] {
				fi = j
			} else if newInterval[0] <= ints[1] {
				fi = j
				newInterval[0] = ints[0]
			}
		}
		if ti == l {
			if newInterval[1] < ints[0] {
				ti = j
			} else if newInterval[1] <= ints[1] {
				ti = j + 1
				newInterval[1] = ints[1]
			}
		}
	}
	//fmt.Println(fi, ti, newInterval)
	//fmt.Println(intervals[:fi], newInterval, intervals[ti:])
	intervals = append(intervals, []int{})
	copy(intervals[fi+1:], intervals[ti:])
	intervals[fi] = newInterval
	return intervals[:len(intervals)-ti+fi]
}

func main() {
	fmt.Println(insert([][]int{}, []int{5, 7}), "=", [][]int{{5, 7}})
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 3}), "=", [][]int{{1, 5}})
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 7}), "=", [][]int{{1, 7}})
	fmt.Println(insert([][]int{{1, 2}}, []int{2, 4}), "=", [][]int{{1, 4}})
	fmt.Println(insert([][]int{{3, 2}}, []int{1, 5}), "=", [][]int{{1, 5}})
	fmt.Println(insert([][]int{{3, 2}}, []int{1, 5}), "=", [][]int{{1, 5}})
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}), "=", [][]int{{1, 5}, {6, 9}})
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}), "=", [][]int{{1, 2}, {3, 10}, {12, 16}})
	fmt.Println(insert([][]int{{1, 5}}, []int{0, 0}), "=", [][]int{{0, 0}, {1, 5}})
	fmt.Println(insert([][]int{{3, 5}, {12, 15}}, []int{6, 6}), "=", [][]int{{3, 5}, {6, 6}, {12, 15}})
}
