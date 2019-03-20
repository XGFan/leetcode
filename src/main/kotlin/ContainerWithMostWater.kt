package com.xulog.algo.deprecated

fun maxArea2(height: IntArray): Int {
    var maxHigh = 0
    var maxArea = 0
    var rHigh = 0
    for (i in 0 until height.size) {
        if (height[i] <= maxHigh) {
            continue
        }
        maxHigh = height[i]
        rHigh = 0
        for (j in (i + 1 until height.size).reversed()) {
            if (height[j] <= rHigh) {
                continue
            }
            rHigh = height[j]
            maxArea = maxOf(minOf(maxHigh, rHigh) * (j - i), maxArea)
        }
    }
    return maxArea
}

fun maxArea(height: IntArray): Int {
    val t = ArrayList<Int>()
    var maxArea = 0
    var maxH = 0
    height.forEachIndexed { index, i ->
        if (i > maxH) {
            maxH = i
            t.add(index)
        }
    }
    maxH = 0
    for (j in (0 until height.size).reversed()) {
        if (height[j] > maxH) {
            maxH = height[j]
            for (i in t) {
                maxArea = maxOf(minOf(maxH, height[i]) * (j - i), maxArea)
            }
        }
    }
    return maxArea
}

fun maxArea3(height: IntArray): Int {
    var maxarea = 0
    var l = 0
    var r = height.size - 1
    while (l < r) {
        maxarea = Math.max(maxarea, Math.min(height[l], height[r]) * (r - l))
        if (height[l] < height[r])
            l++
        else
            r--
    }
    return maxarea
}

fun main(args: Array<String>) {
    val height = intArrayOf(6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 8, 7)
    println(maxArea(height))
    println(maxArea2(height))
    println(maxArea3(height))
}