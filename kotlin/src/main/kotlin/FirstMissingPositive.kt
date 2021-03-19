data class Data constructor(var start: Int, var end: Int) {
    constructor(start: Int) : this(start, start + 1)

    fun moveLeft() {
        this.start = this.start - 1
    }

    fun moveRight() {
        this.end = this.end + 1
    }

    override fun toString(): String {
        return "[$start,$end]"
    }

}

data class Node(
    var data: Data,
    var next: Node? = null
) {
    override fun toString(): String {
        return "$data -> $next"
    }
}

fun firstMissingPositive(nums: IntArray): Int {
    val p = Node(Data(0))
    for (num in nums) {
        if (num > 0 && num <= nums.size) {
            var t: Node? = p
            var tp: Node? = null
            loop@ while (t != null) {
                val tData = t.data
                when {
                    num < tData.start - 1 -> tp?.next = Node(Data(num), t)
                    num == tData.start - 1 -> tData.moveLeft()
                    num == tData.end -> {
                        tData.moveRight()
                        if (t.next?.data?.start == tData.end) {
                            val temp = t.next!!
                            tData.end = temp.data.end
                            t.next = temp.next
                            temp.next = null
                        }
                    }
                    num > tData.end -> {
                        if (t.next == null) {
                            t.next = Node(Data(num))
                        } else {
                            tp = t
                            t = t.next
                            continue@loop
                        }
                    }
                }
                break
            }
        }
    }
    return p.data.end
}


fun main(args: Array<String>) {
    println(firstMissingPositive(arrayOf(1, 2, 0).toIntArray()) == 3)
    println(firstMissingPositive(arrayOf(3, 2).toIntArray()) == 1)
    println(firstMissingPositive(arrayOf(0, -1, 3, 1).toIntArray()) == 2)
    println(firstMissingPositive(arrayOf(4, 1, 2, 3).toIntArray()) == 5)
    println(firstMissingPositive(arrayOf(1, 0, 3, 3, 0, 2).toIntArray()) == 4)
    println(firstMissingPositive(arrayOf(1, 4, 2, 0, 3, 4, 2, 4, 2).toIntArray()) == 5)
    println(firstMissingPositive(arrayOf(-1, 4, 2, 1, 9, 10).toIntArray()) == 3)
    println(firstMissingPositive(arrayOf(2147483647, 2147483646, 2147483645, 3, 2, 1, -1, 0, -2147483648).toIntArray()) == 4)
}








