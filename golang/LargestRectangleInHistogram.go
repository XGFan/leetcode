package main

func largestRectangleArea(heights []int) int {
	max, last := 0, 0
	for i, v := range heights {
		if v == last {
			continue
		} else {
			last = v
		}
		l, r := -1, len(heights)
		for x := i; x >= 0; x-- {
			if heights[x] < v {
				l = x
				break
			}
		}
		for x := i; x < len(heights); x++ {
			if heights[x] < v {
				r = x
				break
			}
		}
		if (r-l-1)*v > max {
			max = (r - l - 1) * v
		}
	}
	return max
}

func main() {

}
