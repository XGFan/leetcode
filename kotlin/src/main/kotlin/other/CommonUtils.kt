package com.xulog.algo.util

/**
 * Created by guofan on 2017/2/18.
 */
fun IntArray.swap(a: Int, b: Int) {
    val temp = this[a]
    this[a] = this[b]
    this[b] = temp
}