fun search(nums: IntArray, target: Int): Int {
    if (nums.isEmpty()) {
        return -1
    }
    if (target == nums[0]) {
        return 0
    }
    val left = target < nums[0]
    var l = 0
    var r = nums.size - 1
    var index = (l + r + 1) / 2 //往r靠
    while (true) {
        if (left) {
            if (nums[index] > nums[0]) {
                if (l == index) {
                    return -1
                }
                l = index
                index = (l + r + 1) / 2
                continue
            }
        } else {
            if (nums[index] < nums[0]) {
                if (r == index) {
                    return -1
                }
                r = index
                index = (l + r + 1) / 2
                continue
            }
        }
        when {
            nums[index] > target -> {
                if (r == index) {
                    return -1
                }
                r = index
                index = (l + r) / 2
            }
            nums[index] < target -> {
                if (l == index) {
                    return -1
                }
                l = index
                index = (l + r + 1) / 2
            }
            else -> return index
        }

    }
}

fun main() {
    println(search(intArrayOf(3, 4, 5, 6, 7, 1, 2, 3), 2))
    println(search(intArrayOf(3, 4, 5, 6, 7, 1), 2))
    println(search(intArrayOf(1), 0))
    println(search(intArrayOf(1, 3), 3))
    println(search(intArrayOf(1, 3), 2))
    println(search(intArrayOf(1, 3), 0))
}

