package com.xulog.algo.sort

class RBTree {

    var p: RBTree? = null
    var v: Int = 0
    var l: RBTree? = null
        set(value) {
            value?.p = this
            field = value
        }
    var r: RBTree? = null
        set(value) {
            value?.p = this
            field = value
        }

    var black: Boolean = false
    val red: Boolean
        get() = !black

    constructor(v: Int, black: Boolean = false) {
        this.v = v
        this.black = black
    }

    companion object {
        @JvmStatic
        fun main(args: Array<String>) {

        }
    }


    fun add(value: Int) {
        if (value <= this.v) {
            if (this.l != null) {
                this.l!!.add(value)
            } else {
                //todo
                if (this.black) { //如果是黑
                    this.l = RBTree(value, false)
                } else {
                    val p = this.p!!
                    if (this == p.r) {
                        //p < value <= v
                        val newTree = RBTree(p.v, true)
                        newTree.l = p.l
                        p.v = value
                        p.l = newTree
                    } else {
                        val newTree = RBTree(p.v, true)
                        newTree.r = p.r
                        p.v = this.v
                        this.v = value
                        p.r = newTree
                    }
                    p.black = false //变成红色
                    this.black = true //变成黑色
                }
            }
        } else {
            if (this.r != null) {
                this.r!!.add(value)
            } else {
                //todo
                if (this.black) { //如果是黑
                    this.r = RBTree(value, false)
                } else {
                    val p = this.p!!
                    if (this == p.l) {
                        val newTree = RBTree(p.v, true)
                        newTree.r = p.r
                        p.v = value
                        p.r = newTree
                    } else {


                    }
                    p.black = false
                    this.black = true

                }
            }
        }

    }

    fun rotateLeft(): RBTree {
        val r = this.r!! //找到右儿子
        this.r = r.l //把右儿子的左键变成自己的右键
        r.l = this //然后把自己变成右儿子的左键
        return r
    }

    fun rotateRight(): RBTree {
        val l = this.l!!
        this.l = l.r //把右儿子的左键变成自己的右键
        l.r = this //然后把自己变成右儿子的左键
        return l
    }


    fun blackCount(pRed: Boolean = false, pCount: Int = 0): Int {
        if (pRed && this.red) {
            throw RuntimeException()
        }
        val thisCount = (if (this.red) 0 else 1) + pCount
        val lc = this.l?.blackCount(this.red, thisCount) ?: thisCount
        val rc = this.r?.blackCount(this.red, thisCount) ?: thisCount
        if (lc != rc) {
            throw RuntimeException()
        } else {
            return lc
        }
    }

    fun isOk(): Boolean {
        try {
            blackCount()
        } catch (e: Exception) {
            e.printStackTrace()
            return false
        }
        return true
    }
}