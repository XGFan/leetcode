fun twoSum(nums: IntArray, target: Int): IntArray {
    val map = hashMapOf<Int, Int>()
    nums.forEachIndexed { index, value ->
        val want = map[target - value]
        if (want != null) {
            return intArrayOf(want, index)
        }
        map[value] = index
    }
    return IntArray(0)
}
