package com.xulog.algo.sort

import com.xulog.algo.util.swap
import java.util.*

/**
 * Created by guofan on 2017/1/22.
 */
object SelectSort {

    fun sort(array: IntArray) {
        for (i in 0..array.size - 1) {
            var flag = i
            for (j in i..array.size - 1) {
                if (array[flag] > array[j]) {
                    flag = j
                }
            }
            if (flag != i) {
                array.swap(i, flag)
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
            SelectSort.sort(unSort)
            assert(Arrays.equals(copy, unSort))
        }
    }

}