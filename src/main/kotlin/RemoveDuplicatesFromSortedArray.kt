fun removeDuplicates(nums: IntArray): Int {
    var count = -1
    for (num in nums) {
        if (nums.getOrNull(count) != num) {
            count++
            nums[count] = num
        }
    }
    return count + 1
}