fun searchInsert(nums: IntArray, target: Int): Int {
    var l = 0
    var r = nums.size - 1
    var p = (nums.size - 1) / 2
    while (l <= r) {
        when {
            nums[p] < target -> l = p + 1 //l指向p右边一点
            nums[p] > target -> r = p - 1 //r指向p左边一点
            else -> return p
        }
        p = (l + r) / 2 //总是偏向左边一点
    }
    return l
}

fun main() {
    println(searchInsert(intArrayOf(1, 3, 5, 6), 5))
    println(searchInsert(intArrayOf(1, 3, 5, 6), 2))
    println(searchInsert(intArrayOf(1, 3, 5, 6), 7))
    println(searchInsert(intArrayOf(1, 3, 5, 6), 0))
}
