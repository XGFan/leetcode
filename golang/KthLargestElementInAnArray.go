package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type List []int

func quickSort(arrays []int) {
	quickSortByRange(arrays, 0, len(arrays)-1)
}

func quickSortByRange(arrays []int, start, end int) {
	if end > start {
		//确定轴点
		pivot := partitionList(arrays, start, end)
		quickSortByRange(arrays, start, pivot-1)
		quickSortByRange(arrays, pivot+1, end)
	}
}

func findKthLargest2(nums []int, k int) int {
	quickSort(nums[:k])
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			index := k - 1
			for j := 1; j < k; j++ {
				if nums[j] >= nums[i] {
					index = j - 1
					break
				}
			}
			copy(nums[:index], nums[1:index+1])
			nums[index] = nums[i]
		}
	}
	return nums[0]
	//return findKthLargestByRange(nums, 0, len(nums)-1, k)
}

//func sortLast(list []int) {
//	for i := len(list) - 2; i >= 0; i-- {
//		if list[i] < list[i+1] {
//			list[i], list[i+1] = list[i+1], list[i]
//		}
//	}
//}

func partitionList(arrays []int, start, end int) int {
	l := start
	pivot := arrays[end]
	for p := start; p < end; p++ {
		if arrays[p] < pivot {
			arrays[p], arrays[l] = arrays[l], arrays[p]
			l++
		}
	}
	arrays[l], arrays[end] = arrays[end], arrays[l]
	return l
}
func findKthLargestByRange(nums []int, start, end, k int) int {
	if end > start {
		//确定轴点
		pivot := partitionList(nums, start, end)
		if pivot == k {
			return nums[k]
		} else if pivot > k {
			return findKthLargestByRange(nums, start, pivot-1, k)
		} else {
			return findKthLargestByRange(nums, pivot+1, end, k)
		}
	}
	return nums[start]
}
func findKthLargest(nums []int, k int) int {
	return findKthLargestByRange(nums, 0, len(nums)-1, k)
}

//func findKthLargest(nums []int, k int) int {
//	list := make([]int, 0, k)
//	for i, num := range nums {
//		if i < k {
//			list = append(list, num)
//			sortLast(list)
//		} else {
//			if list[k-1] < num {
//				list[k-1] = num
//				sortLast(list)
//			}
//		}
//	}
//	return list[k-1]
//}

func sortLast(list []int) {
	for i := len(list) - 2; i >= 0; i-- {
		if list[i] < list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}
	}
}

//func (list *Node) insertToNums(num int) *Node {
//	newNode := &Node{Value: num}
//	if list == nil {
//		return newNode
//	} else {
//		if list.Value >= num {
//			newNode.Next = list
//			return newNode
//		}
//		prev, next := list, list.Next
//		for next != nil {
//			if next.Value >= num {
//				prev.Next, newNode.Next = newNode, next
//				return list
//			} else {
//				prev, next = next, next.Next
//			}
//		}
//		prev.Next = newNode
//		return list
//	}
//}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	ints := []int{3, 2, 1, 5, 6, 4}
	quickSort(ints)
	fmt.Println(ints)
}
