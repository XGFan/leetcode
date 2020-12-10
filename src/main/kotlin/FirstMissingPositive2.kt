fun IntArray.swap(a: Int, b: Int) {
    val temp = this[a]
    this[a] = this[b]
    this[b] = temp
}

fun firstMissingPositive2(nums: IntArray): Int {
    if (nums.isEmpty()) {
        return 1
    }
    var index = 0
    while (index < nums.size) {
        val value = nums[index]
        val orderShouldBe = value - 1
        //正数 && 小于等于长度 && 不在应该在的位置 && 应该在的位置不是现在的值
        if (value > 0 && value <= nums.size && index != orderShouldBe && nums[orderShouldBe] != value) {//往前换或者往后换
            nums.swap(index, orderShouldBe)
        } else {
            index++
        }
    }
    if (nums[0] != 1)
        return 1
    index = 0
    while (index <= nums.size - 1 && nums[index] == index + 1) {
        index++
    }
    return index + 1
}

fun main(args: Array<String>) {
    println(firstMissingPositive2(arrayOf(0, 2, 2, 1, 1, 0).toIntArray()) == 3)
    println(firstMissingPositive2(arrayOf(1, 3, 3).toIntArray()) == 2)
    println(firstMissingPositive2(arrayOf(2, 2).toIntArray()) == 1)
    println(firstMissingPositive2(arrayOf(1, 2, 0).toIntArray()) == 3)
    println(firstMissingPositive2(arrayOf(3, 2).toIntArray()) == 1)
    println(firstMissingPositive2(arrayOf(0, -1, 3, 1).toIntArray()) == 2)
    println(firstMissingPositive2(arrayOf(4, 1, 2, 3).toIntArray()) == 5)
    println(firstMissingPositive2(arrayOf(3, 4, -1, 1).toIntArray()) == 2)
    println(firstMissingPositive2(arrayOf(1, 0, 3, 3, 0, 2).toIntArray()) == 4)
    println(firstMissingPositive2(arrayOf(1, 4, 2, 0, 3, 4, 2, 4, 2).toIntArray()) == 5)
    println(firstMissingPositive2(arrayOf(-1, 4, 2, 1, 9, 10).toIntArray()) == 3)
    println(firstMissingPositive2(arrayOf(2, 4, 1, 5, 6).toIntArray()) == 3)
    println(
        firstMissingPositive2(
            arrayOf(
                2147483647,
                2147483646,
                2147483645,
                3,
                2,
                1,
                -1,
                0,
                -2147483648
            ).toIntArray()
        ) == 4
    )
}