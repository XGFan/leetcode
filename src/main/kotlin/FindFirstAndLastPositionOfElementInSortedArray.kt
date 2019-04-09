import kotlin.random.Random

fun searchRange(nums: IntArray, target: Int): IntArray {
    var l = 0
    var r = nums.size - 1
    var p = (nums.size - 1) / 2
    while (l <= r) {
        when {
            nums[p] < target -> l = p + 1 //l指向p右边一点
            nums[p] > target -> r = p - 1 //r指向p左边一点
            else -> return intArrayOf(search(nums, l, p, target, true), search(nums, p, r, target, false))
        }
        p = (l + r) / 2 //总是偏向左边一点
    }
    return intArrayOf(-1, -1)
}


fun search(nums: IntArray, l: Int, r: Int, target: Int, left: Boolean): Int {
    var l = l
    var r = r
    var p = if (left) {
        (r + l) / 2 //如果在左边比较，需要尽可能的靠近左边
    } else {
        (r + l + 1) / 2
    }
    while (l < r) {
        when {
            nums[p] < target -> l = p + 1
            nums[p] > target -> r = p - 1
            else -> {
//                println("${if (left) "left" else "right"} $l $p $r")
                if (left) {
                    if (l != p) {
                        r = p //r指向新边界,但是不能抛弃掉这一点
                    } else {
                        return l
                    }
                } else {
                    if (r != p) {
                        l = p
                    } else {
                        return p
                    }
                }
            }
        }
        p = if (left) {
            (r + l) / 2 //如果在左边比较，需要尽可能的靠近左边
        } else {
            (r + l + 1) / 2
        }
    }
    return p
}


fun main() {
    repeat(100000) {
        val list = List(Random.nextInt(50, 100)) {
            Random.nextInt(10, 100)
        }.sorted()
        val target = Random.nextInt(10, 100)
        if (!intArrayOf(list.indexOfFirst { it == target }, list.indexOfLast { it == target })
                .contentEquals(searchRange(list.toIntArray(), target))
        ) {
            println("!!!")
        }

    }
//    println(searchRange(intArrayOf(), 8).contentToString())
//    println(searchRange(intArrayOf(5, 7, 7, 8, 8, 10), 8).contentToString())
//    println(searchRange(intArrayOf(5, 7, 7, 8, 8, 10), 6).contentToString())
//    println(searchRange(intArrayOf(1, 3), 1).contentToString())
//    println(searchRange(intArrayOf(1, 4), 4).contentToString())
//    println(searchRange(intArrayOf(2, 2), 2).contentToString())
}