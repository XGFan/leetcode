fun removeElement(nums: IntArray, `val`: Int): Int {
    var count = 0
    for (num in nums) {
        if (`val` != num) {
            nums[count] = num
            count++
        }
    }
    return count
}