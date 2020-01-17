package com.xulog.algo.sort

import com.xulog.algo.util.swap
import java.util.*

/**
 * Created by guofan on 2017/1/22.
 */
object BubbleSort {

    fun sort(array: IntArray) {
        for (i in 0..array.size - 1) {
            for (j in i + 1..array.size - 1) {
                if (array[i] > array[j]) {
                    array.swap(i, j)
                }
            }
        }
    }


    @JvmStatic
    fun main(args: Array<String>) {
        (1..10000).forEach {
            val unSort = (1..30).map {
                Random().nextInt(100)
            }.toTypedArray().toIntArray()
            val copy = unSort.clone()
            copy.sort()
            BubbleSort.sort(unSort)
            assert(Arrays.equals(copy, unSort))
        }
    }

}