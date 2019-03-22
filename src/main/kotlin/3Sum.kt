fun threeSum(nums: IntArray): List<List<Int>> {
    val arrayList = ArrayList<List<Int>>()
    nums.sort()
    for (i in 0 until nums.size) {
        if (i > 0 && nums[i] == nums[i - 1]) {
            continue
        }
        val want = 0 - nums[i]
        var l = i + 1
        var r = nums.size - 1
        while (l < r && l >= i + 1 && r <= nums.size - 1) {
            if (l > i + 1 && nums[l] == nums[l - 1]) {
                l++
                continue
            }
            if (r < nums.size - 1 && nums[r] == nums[r + 1]) {
                r--
                continue
            }
            when {
                nums[l] + nums[r] < want -> l++
                nums[l] + nums[r] == want -> {
                    arrayList.add(listOf(nums[i], nums[l], nums[r]))
                    l++
                    r--
                }
                nums[l] + nums[r] > want -> r--
            }
        }
    }
    return arrayList
}
