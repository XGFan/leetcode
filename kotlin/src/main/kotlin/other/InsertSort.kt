package com.xulog.algo.sort

import java.util.*

/**
 * Created by guofan on 2017/1/22.
 */
object InsertSort {
    fun <T : Comparable<T>> sort(list: Collection<T>): List<T> {
        val linkedList = LinkedList(list)
        for (i in 1..linkedList.size - 1) {
            for (j in 0..i - 1) {
                if (linkedList[i] < linkedList[j] && (j == 0 || linkedList[i] >= linkedList[j - 1])) {
                    linkedList.add(j, linkedList[i])
                    linkedList.removeAt(i + 1)
                }
            }
        }
        return linkedList
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
