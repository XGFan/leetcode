package other

import kotlin.math.min

/**
 * Created by Guofan on 2017/1/24.
 */
object Merge {

    class Node(var value: Int) {
        var next: Node? = null

        override fun toString(): String {
            return if (next == null) {
                "$value"
            } else {
                "$value,$next"
            }
        }
    }


    fun merge(intArray: IntArray): IntArray {
        var i = 1
        var array = intArray
        while (i < array.size ) {
            var j = 0
            while (j +i  < array.size ) {
                array = mergeTwoLists(array, j, i, j + i, min(array.size - j - i, i))
                println(array.toList())
                j += i
            }
            i *= 2
        }
        return array
    }

    fun mergeTwoLists(array: IntArray, x: Int, xl: Int, y: Int, yl: Int): IntArray {
        val newArray = IntArray(array.size)
        var index = x
        var i = 0
        var j = 0
        while (i < xl && j < yl) {
            if (array[x + i] < array[y + j]) {
                newArray[index] = array[x + i]
                i++
            } else {
                newArray[index] = array[y + j]
                j++
            }
            index++
        }
        if (i == xl) {
            System.arraycopy(array, y + j, newArray, y + j, array.size - y - j)
        } else {
            System.arraycopy(array, x + i, newArray, x + i, array.size - x - i)
        }
        System.arraycopy(array, y + yl, newArray, y + yl, array.size - y - yl)
        return newArray
    }

    @JvmStatic
    fun main(args: Array<String>) {
        println(merge(intArrayOf(6, 2, 4, 3, 5, 1)).toList())

//        (1..10000).forEach {
//            val unSort = (1..22).map {
//                Random().nextInt(100)
//            }.toList()
//            assert(unSort.sorted() == merge(unSort.toTypedArray().toIntArray()).toList())
//        }
    }
}