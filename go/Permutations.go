package main

func permuteIter(result *[][]int, prefix, nums []int, skip uint8) {
	if len(prefix) == len(nums) {
		*result = append(*result, make([]int, len(prefix)))
		copy((*result)[len(*result)-1], prefix)
	} else {
		for i, v := range nums {
			x := uint8(1 << i)
			if skip&x == 0 {
				permuteIter(result, append(prefix, v), nums, skip|x)
			}
		}
	}
}

func permute(nums []int) [][]int {
	var size = 1
	for i := range nums {
		size *= 1 + i
	}
	result := make([][]int, 0, size)
	permuteIter(&result, make([]int, 0, len(nums)), nums, 0)
	return result
}

func main() {

}
