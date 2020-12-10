import com.xulog.algo.util.swap
import java.util.*

/**
 * Created by guofan on 2017/2/13.
 */
object HeapSort {

    tailrec fun casWithFather(array: IntArray, index: Int) {
        if (index == 0) {
            //第一个节点 do nothing
        } else {
            //和父节点比，如果大于父节点，递归和父节点比。
            val parent = (index + 1) / 2 - 1
            if (array[parent] < array[index]) {
                array.swap(index, parent)
                casWithFather(array, parent)
            }
        }
    }

    tailrec fun casWithChild(array: IntArray, farther: Int, end: Int) {
        val lc = farther * 2 + 1
        val rc = farther * 2 + 2
        if (lc >= end) {
            return
        }
        val better = if (array[lc] > array[rc] || rc >= end) {
            lc
        } else {
            rc
        }
        if (array[farther] < array[better]) {
            array.swap(farther, better)
            casWithChild(array, better, end)
        }
    }

    fun sort(array: IntArray) {
        for (i in 0..array.size - 1) {
            casWithFather(array, i)
        }
        for (i in (1..array.size - 1).reversed()) {
            array.swap(0, i)
            casWithChild(array, 0, i)
        }
    }


    @JvmStatic
    fun main(args: Array<String>) {
        val arraysList = (1..10000).map {
            ((1..500).map {
                Random().nextInt(100)
            }).toTypedArray().toIntArray()
        }
        val arraysList2 = arraysList.map(IntArray::clone)

        var start = System.currentTimeMillis()
        arraysList.forEach { HeapSort.sort(it) }
        var end = System.currentTimeMillis()
        println("My sort spent:${end - start}")

        start = System.currentTimeMillis()
        arraysList2.forEach(IntArray::sort)
        end = System.currentTimeMillis()
        println("std sort spent:${end - start}")

        arraysList.forEachIndexed { i, ints ->
            if (!Arrays.equals(ints, arraysList2[i])) {
                println("fuck")
            }
        }

    }
}