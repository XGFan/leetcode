package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	l := (len1 + len2 - 1) / 2
	r := (len1 + len2) / 2
	var lp, rp int //left point & right point
	var p1, p2 int //list1 point, list2 point
	for i := 0; i <= r; i++ {
		lp = rp
		if p1 == len1 {
			//reach end of list1
			if l >= i { //l still need to move
				lp = nums2[p2+l-i]
			}
			rp = nums2[p2+r-i]
			break
		}
		if p2 == len2 {
			//reach end of list2
			if l >= i { //l still need to move
				lp = nums1[p1+l-i]
			}
			rp = nums1[p1+r-i]
			break
		}
		if nums1[p1] < nums2[p2] {
			rp = nums1[p1]
			p1++
		} else {
			rp = nums2[p2]
			p2++
		}
	}
	if r == l {
		return float64(rp)
	} else {
		return float64(lp+rp) / 2
	}
}

func main() {
	println(findMedianSortedArrays([]int{2}, []int{1, 3, 4})) //0,0
	//println(findMedianSortedArrays([]int{2, 3}, []int{}))                 //0,0
	//println(findMedianSortedArrays([]int{}, []int{1}))                    //0,0
	//println(findMedianSortedArrays([]int{2}, []int{}))                    //0,0
	//println(findMedianSortedArrays([]int{}, []int{1, 2, 3, 4}))           //1,2
	//println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))             //1,2
	//println(findMedianSortedArrays([]int{1, 3, 5, 7}, []int{2, 4, 6, 8})) //1,2
	//println(findMedianSortedArrays([]int{2}, []int{1, 3, 4}))             //1,2
}
