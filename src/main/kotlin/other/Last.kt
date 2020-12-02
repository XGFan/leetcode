package other

import com.xulog.algo.util.swap

object Last {
    fun f(m: Long, n: Long): Long {
        if (m == 0L && n == 0L) {
            return 1L
        }
        if (m < 0L || n < 0L) {
            return 1L
        }
        if (n == 0L) {
            return 3L
        }
        var total = 0L
        for (i in 0 until n) {
            total += f(m - 1, i)
        }
        return n + f(m, 0L) + total
    }


    fun fff(prefix: String, chars: List<Char>): List<String> {
        if (chars.isEmpty()) {
            return listOf(prefix)
        }
        return chars.flatMap {
            fff(prefix + it, (chars - it))
        }
    }

    @JvmStatic
    fun main(args: Array<String>) {
        val charRange = '1'..'5'
        val solution = Solution()
//        println(
        fff("", charRange.toList())
            .sorted()
            .map { it -> it.map { it.toString().toInt() }.toIntArray() }
//            .flatMap { Arrays.asList(it.toInt(),it.toInt()) }
            .forEach { solution.nextPermutation(it) }
//            .map { it.toInt() }
//            .joinToString("\n"))
//        println(f(10L, 20L))
//        println(f(50L, 100L))

//        println(listOf(1, 3, 1, 4, 5, 9, 7, 1).toIntArray().last(1).toInt())
//        println(listOf(1, 3, 1, 4, 5, 9, 7, 1).toIntArray().last(2).toInt())
    }


    fun IntArray.last(i: Int): IntArray {
        return List(i) {
            this[this.size - it - 1]
        }
            .reversed()
            .toIntArray()
    }

    fun IntArray.toInt(): Int {
        return this.fold(0) { acc, v -> acc * 10 + v }
    }


    class Solution {
        fun nextPermutation(nums: IntArray): IntArray {
            var p = 0
            for (i in nums.indices.reversed()) {
                if ((i - 1 < 0) || nums[i] > nums[i - 1]) {
                    p = i
                    break
                }
            }
            print("${nums.toInt()} -> ")
//            print("${nums.last(nums.size - x + 1).toInt()}")
            var j = nums.size - 1
            if (p != 0) {
                val p = p - 1 //需要更换
                for (i in p + 1 until nums.size) {
                    if (nums[i] <= nums[p - 1]) {
                        j = i - 1
                        break
                    }
                }
                nums.swap(p - 1, j)
                j = nums.size - 1
            }

            while (p < j) {
                nums.swap(p, j)
                p++
                j--
            }
            print("${nums.toInt()}")
            println()
            return nums
//            print("\t${nums[x - 1]}:${nums[j]}")


        }

//        fun nextPermutation(nums: IntArray): IntArray {
//            for (i in nums.indices) {
//                if ((nums.getOrNull(i + 1) == null || nums[i] > nums[i + 1]) &&
//                    (nums.getOrNull(i - 1) == null) || nums[i] > nums[i - 1]
//                ) {
//                    println(i)
//                }
//            }
//            return nums
//        }
    }
}