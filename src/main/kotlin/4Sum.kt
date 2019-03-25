fun fourSum(nums: IntArray, target: Int): List<List<Int>> {
    val arrayList = ArrayList<List<Int>>()
    nums.sort()
    for (i in 0 until nums.size) {
        if (i > 0 && nums[i] == nums[i - 1]) {
            continue
        }
        for (j in (i until nums.size).reversed()) {
            if (j < nums.size - 1 && nums[j] == nums[j + 1]) {
                continue
            }
            val want = target - nums[i] - nums[j]
            var l = i + 1
            var r = j - 1
            while (l < r && l >= i + 1 && r <= j - 1) {
                if (l > i + 1 && nums[l] == nums[l - 1]) {
                    l++
                    continue
                }
                if (r < j - 1 && nums[r] == nums[r + 1]) {
                    r--
                    continue
                }
                when {
                    nums[l] + nums[r] < want -> l++
                    nums[l] + nums[r] == want -> {
                        arrayList.add(listOf(nums[i], nums[l], nums[r], nums[j]))
                        l++
                        r--
                    }
                    nums[l] + nums[r] > want -> r--
                }
            }
        }
    }
    return arrayList
}

fun main() {
    println(fourSum(intArrayOf(0, 0, 0, 0), 0))
}