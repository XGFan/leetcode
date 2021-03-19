package com.xulog.algo.sort

import java.util.*

/**
 * Created by guofan on 2017/1/22.
 */
object Select {

    fun <T : Comparable<T>> sort(list: Collection<T>): List<T> {
        val arrayList = ArrayList(list)
        for (i in 0..arrayList.size - 1) {
            var flag = i
            for (j in i..arrayList.size - 1) {
                if (arrayList[flag] > arrayList[j]) {
                    flag = j
                }
            }
            if (flag != i) {
                val temp = arrayList[i]
                arrayList[i] = arrayList[flag]
                arrayList[flag] = temp
            }
        }
        return arrayList
    }


    @JvmStatic
    fun main(args: Array<String>) {
        (1..10000).forEach {
            val unSort = (1..20).map {
                Random().nextInt(100)
            }.toList()
            assert(unSort.sorted() == sort(unSort))
        }
    }

}