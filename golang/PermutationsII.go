package main

func permuteUniqueIter(result *[][]int, prefix, nums []int, skip int) {
	// fmt.Printf("%v %v \n", prefix, result)
	if len(prefix) == len(nums) {
		*result = append(*result, make([]int, len(prefix)))
		copy((*result)[len(*result)-1], prefix)
	} else {
		ignore := 0
		for i, v := range nums {
			x := 1 << i
			if skip&x == 0 {
				y := 1 << (v + 10)
				if ignore&y == 0 {
					permuteUniqueIter(result, append(prefix, v), nums, skip|x)
					ignore |= y
				}
			}
		}
	}
}
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0, 1)
	permuteUniqueIter(&result, make([]int, 0, len(nums)), nums, 0)
	return result
}

func main() {

}
