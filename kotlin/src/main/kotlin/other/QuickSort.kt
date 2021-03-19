package com.xulog.algo.sort

import com.xulog.algo.util.swap
import java.util.*

/**
 * Created by guofan on 2017/2/18.
 */
object QuickSort {
    @JvmOverloads
    fun sort(arrays: IntArray, start: Int = 0, end: Int = arrays.size - 1) {
        if (end > start) {
            //确定轴点
            val pivot = partition(arrays, start, end)
            sort(arrays, start, pivot - 1)
            sort(arrays, pivot + 1, end)
        }
    }

    fun partition(arrays: IntArray, start: Int, end: Int): Int {
        var l = start
        val pivot = arrays[end]
        for (p in start..end - 1) {
            if (arrays[p] < pivot) {
                if (l != p) {
                    arrays.swap(p, l)
                }
                l++
            }
        }
        arrays.swap(l, end)
        return l
    }

    @JvmStatic
    fun main(args: Array<String>) {
        (1..20).forEach {
            val unSort = (1..30).map {
                Random().nextInt(100)
            }.toTypedArray().toIntArray()
            val copy = unSort.clone()
            copy.sort()
            QuickSort.sort(unSort)
            assert(Arrays.equals(copy, unSort))
        }
    }
}