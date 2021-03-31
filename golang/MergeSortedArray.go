package main

func merge_(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m+n]
	i, j := 0, 0
	for {
		if i >= m {
			copy(nums1[len(nums1)-n+j:], nums2[j:n])
			return
		}
		if j >= n {
			return
		}
		if nums1[i] > nums2[j] {
			copy(nums1[i+1:], nums1[i:])
			nums1[i] = nums2[j]
			m++
			j++
		}
		i++
	}
}

func main() {
	//ints := []int{1, 7, 9, 11}
}
