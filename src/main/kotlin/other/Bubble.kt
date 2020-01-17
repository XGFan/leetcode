package com.xulog.algo.sort

import java.util.*
import kotlin.system.measureNanoTime

/**
 * Created by guofan on 2017/1/22.
 */
object Bubble {

    fun <T : Comparable<T>> sort(list: Collection<T>): List<T> {
        val arrayList = ArrayList(list)
        for (i in 0..arrayList.size - 1) {
            for (j in Math.min(i + 1, arrayList.size - 1)..arrayList.size - 1) {
                if (arrayList[i] > arrayList[j]) {
                    //swap it
                    val temp = arrayList[i]
                    arrayList[i] = arrayList[j]
                    arrayList[j] = temp
                }
            }
        }
        return arrayList
    }

    @JvmStatic
    fun main(args: Array<String>) {
        measureNanoTime {
            repeat(10000) {
                val unSort = (1..20).map {
                    Random().nextInt(100)
                }.toList()
                assert(unSort.sorted() == sort(unSort))
            }
        }
    }
}