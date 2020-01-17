package com.xulog.algo.sort

import java.util.*

class T {


    companion object {
        fun randomTest() {
            List(1000) { index ->
                repeat(1000) {
                    val gTree = Tree(10, null)
                    val linkedList = LinkedList<Int>()
                    linkedList.add(10)
                    for (i in 1..index) {
                        val v = Random().nextInt(30)
                        gTree.addTree(Tree(v, null))
                        linkedList.add(v)
                    }
                    val copy = linkedList.toList()
                    linkedList.sort()
                    while (true) {
                        val delMin = gTree.delMinLeafToRoot()
                        val test = linkedList.pollFirst()
                        if (delMin != test) {
                            println("fuckyou")
                            break
                        }
                        if (delMin == null) {
                            break
                        }
                    }
                }
            }
        }

        fun test(case: List<Int>) {
            val list = LinkedList(case)
//            list.shuffle()
            val pinned = list.toList()
            val first = list.pollFirst()
            val gTree = Tree(first, null)
            val gTree2 = Tree(first, null)
            for (i in list) {
                gTree.addTree(Tree(i, null))
                gTree2.addTree(Tree(i, null))
            }
            val sorted = pinned.sorted()
            for (i in sorted) {
                val delMin = gTree.delMinRootToLeaf()
                val delMin2 = gTree2.delMinLeafToRoot()
                if (delMin != delMin2) {
                    println("$delMin != $delMin2")
                } else {
                    println("$delMin == $delMin2")
                }
            }
        }

        fun testRemove(case: List<Int> = emptyList()) {
            val list = LinkedList(case)
            if (list.isEmpty()) {
                repeat(Random().nextInt(30) + 30) {
                    list.add(Random().nextInt(50))
                }
            }
            val pinned = list.toList()
            val gTree = Tree(list.first, null)
            list.drop(1).forEach {
                gTree.addTree(Tree(it, null))
            }
            for (i in list) {
                gTree.remove(i)
                if (!gTree.isOK()) {
                    println(pinned)
                    System.exit(-1)
                }
            }
        }

        fun generateList(len: Int = 29): LinkedList<Int> {
            val random = Random()
            val list = LinkedList<Int>()
            repeat(len) {
                list.add(random.nextInt(20))
            }
            return list
        }


        fun randomTestAdd() {
            while (true) {
                val sb = StringBuffer()
//                val list = generateList(9)
                val list = listOf(2, 2, 17, 1, 2)
                val tree = Tree(5)
                list.forEach {
                    tree.add(it)
                    sb.append(it).append(",")
                    if (!tree.isOK()) {
                        sb.appendln()
                        sb.append("[MAX] Generate Tree Fail ${tree}").appendln()
                        println(sb)
                        System.exit(0)
                    }
                }
            }
        }


        fun randomTestMax() {
            while (true) {
                val sb = StringBuffer()
                val list = generateList(9)

                val tree = Tree(5)
                list.forEach {
                    tree.add(it)
                    if (!tree.isOK()) {
                        sb.append("[MAX] Generate Tree Fail ${tree}").appendln()
                        println(sb)
                        System.exit(0)
                    }
                }
                list.addFirst(5)

                val pinned = list.toList()
                sb.append(pinned).appendln()

                var max = list.max()
                while (max != null) {
                    list.remove(max)
                    sb.append("[MAX] Tree Before Delete ${tree}").appendln()
                    val treeMax = tree.delMaxLeafToRoot()
                    if (!tree.isOK()) {
                        sb.append("[MAX] Tree After Delete ${tree}").appendln()
                        println(sb)
                        System.exit(0)
                    }
                    if (treeMax != max) {
                        sb.append("[MAX] $max != $treeMax").appendln()
                        println(sb)
                        System.exit(0)
                    } else {
                        sb.append("[MAX] $max == $treeMax").appendln()
                    }
                    max = list.max()
                }
            }

        }


        fun randomTestAll() {
            val list = LinkedList(listOf(50))
            val tree = Tree(50)
            var addOp = 0
            var delOp = 0
            var mOp = 0

            val random = Random()
            loop@ while (true) {
                val op = random.nextInt(85)

                val pinnedList = list.sorted()
                val pinnedTree = tree.toString()
                when (op) {
                    in 0..25 -> {
                        val value = list.pollFirst()
                        if (value != null) {
                            try {
                                tree.remove(value)
                                if (!tree.isOK()) {
                                    println("[REMOVE] $value")
                                    System.currentTimeMillis()
                                }
                            } catch (e: Exception) {
                                println("[REMOVE] FAIL")
                                e.printStackTrace()
                                break@loop
                            }
                            delOp++
                        }
                    }
                    in 25..30 -> { //getAndRemoveMin
                        val min = list.min()
                        if (min != null) {
                            list.remove(min)
                            val treeMin = tree.delMinLeafToRoot()
                            if (!tree.isOK()) {
                                println(pinnedTree)
                                System.exit(-1)
                            }
                            if (treeMin != min) {
                                println("[MIN] $min != $treeMin")
                                break@loop
                            }
                            if (!tree.isOK()) {
                                println("[MIN] $treeMin")
                                println(pinnedTree)
                                System.exit(-1)
                            }
                            mOp++
                        }
                    }
                    in 30..35 -> { //getAndRemoveMin
                        val max = list.max()
                        if (max != null) {
                            list.remove(max)
                            val treeMax = tree.delMaxLeafToRoot()
                            if (treeMax != max) {
                                println("[MAX] $max != $treeMax")
                                break@loop
                            }
                            if (!tree.isOK()) {
                                println("[MAX] $treeMax")
                                println(pinnedTree)
                                System.exit(-1)
                            }
                            mOp++
                        }
                    }
                    else -> {
                        if (list.size >= 1_0000) {
                            break@loop
                        } else {
                            val value = random.nextInt(100)
                            list.add(value)
                            tree.add(value)
                            addOp++
                            if (!tree.isOK()) {
                                println("[ADD] $value")
                                println(pinnedTree)
                                System.exit(-1)
                            }
                        }
                    }
                }
            }
            println("ADD:$addOp")
            println("MIN&MAX:$mOp")
            println("DELETE:$delOp")
        }

        @JvmStatic
        fun main(args: Array<String>) {
            randomTestAll()
        }
    }
}


fun Tree.is3Tree(): Boolean {
    return this.v.size == 2
}

fun Tree.is2Tree(): Boolean {
    return this.v.size == 1
}


class Tree : Comparable<Tree> {
    var p: Tree? = null
    var v: LinkedList<Int> = LinkedList()
    var c: LinkedList<Tree> = LinkedList()

    val type
        get() = "${v.size}-${c.size}"

    constructor()

    constructor(v: Int, p: Tree? = null) {
        this.p = p
        this.v.add(v)
    }

    constructor(str: String) {
        var p = Tree()
        var temp = 0
        for (c in str) {
            when (c) {
                '[' -> {
                    p.addChild(Tree())
                    p = p.c.last
                }
                '(' -> {
                }
                ')' -> {
                    p.v.add(temp)
                    temp = 0
                }
                ']' -> {
                    p = p.p!!
                }
                else -> {
                    temp = temp * 10 + c.toInt() - 48
                }
            }
        }
        this.c.clear()
        this.autoCombine(p.c.first)
    }

    constructor(nv: Int, l: Tree, r: Tree, p: Tree?) : this(nv, p) {
        l.p = this
        r.p = this
        this.addChild(l)
        this.addChild(r)
    }

    fun addFirstChild(child: Tree) {
        child.p = this
        this.c.addFirst(child)
    }

    fun addLastChild(child: Tree) {
        child.p = this
        this.c.addLast(child)
    }

    fun addChild(child: Tree) {
        child.p = this
        this.c.add(child)
    }

    fun addChild(index: Int, child: Tree) {
        child.p = this
        this.c.add(index, child)
    }

    fun upgrade(tree: Tree) {
        val index = this.c.indexOf(tree)
        this.c.remove(tree)
        this.addTree(tree, index)
    }


    private fun getIndexInChild(value: Int): Int {
        var index: Int = this.v.size - 1
        while (index >= 0) {
            if (this.v[index] < value) {
                break
            } else {
                index--
            }
        }
        return index + 1
    }


    fun autoCombine(tree: Tree) {
        this.v.addAll(tree.v)
        this.v.sort()
        tree.c.forEach {
            it.p = this
            this.c.add(it)
        }
        this.c.sort()
    }


    fun addTree(tree: Tree, childIndex: Int = -1) {
        when (tree.type) {
            "1-0" -> {//添加一个单点
                when (this.type) {
                    "0-0" -> this.v.add(tree.v[0])
                    "1-0" -> this.v.add(getIndexInChild(tree.v[0]), tree.v[0]) //转变成一个 empty-3-tree
                    "2-0" -> { //转变成一个新的full-2-tree
                        this.v.add(getIndexInChild(tree.v[0]), tree.v[0])
                        this.addChild(Tree(this.v.pollFirst(), this))
                        this.addChild(Tree(this.v.pollLast(), this))
//                        if (this.p != null) {
//                            val index = this.p!!.c.indexOf(this)
//                            this.p!!.unLink(this).addTree(this, index)
//                        }
                        this.p?.upgrade(this)
                    }
                    "1-2", "2-3", "3-4" -> this.c[getIndexInChild(tree.v[0])].addTree(tree)  //full-tree存在子节点，让子节点去处理
                    else -> throw RuntimeException("状态不正确，添加${tree.type}到${this.type}")
                }
            }
            "1-2" -> {//添加一个full-2-tree
//                if (childIndex == -1) {
//                    this.autoCombine(tree)
//                } else {
                this.v.add(childIndex, tree.v[0])
                for (t in tree.c.reversed()) {
                    this.addChild(childIndex, t)
                }
//                }
                when (this.type) {
                    "2-3" -> {//full-2-tree（某个子节点进位) -> full-3-tree -> keep
                    }
                    "3-4" -> {//full-3-tree（某个子节点进位) -> full-4-tree -> 2 full-2-tree
                        val l = Tree(this.v.pollFirst(), this.c.pollFirst(), this.c.pollFirst(), this)
                        val r = Tree(this.v.pollLast(), this.c.pollFirst(), this.c.pollFirst(), this)
                        this.addChild(l)
                        this.addChild(r)
//                        this.p?.unLink(this)?.addTree(this)
                        this.p?.upgrade(this)
                    }
                    "4-5" -> {//这tm一般都是正在规整中的根节点
                        //经过上上面两步，这已经是一个4-5的神奇节点了
                        val lTree = Tree(this.v.pollFirst(), null)//拿出最左侧的值
                        lTree.v.add(this.v.pollFirst()) //再加上左二，成为了一个 2-0节点
                        repeat(3) {
                            //成为了一个2-3节点
                            lTree.addChild(this.c.pollFirst())
                        }
                        val rTree = Tree(this.v.pollLast(), null) //再生成右树
                        repeat(2) {
                            rTree.addChild(this.c.pollFirst()) //1-2
                        }
                        this.addChild(lTree)
                        this.addChild(rTree)
//                        this.p?.unLink(this)?.addTree(this)
                        this.p?.upgrade(this)
                    }
                    else -> throw RuntimeException("状态不正确，添加${tree.type}到${this.type}")
                }
            }
            else -> throw RuntimeException("状态不正确，添加${tree.type}到${this.type}")
        }
    }

    fun add(value: Int) {
        this.addTree(Tree((value)))
    }


    fun remove(value: Int): Boolean {
        val childIndex = getIndexInChild(value)
        return if (childIndex < this.v.size && value == this.v[childIndex]) {
            //todo 找到了
            if (this.c.isEmpty()) {
                this.transform() //todo ugly
                if (this.v.isEmpty()) {
                    System.currentTimeMillis() //冷静一下
                    this.p!!.remove(value)
                } else {
                    val remove = this.v.remove(value)
                    remove
                }
            } else {
                val maxTreeNode = this.c[childIndex].maxTreeNode() //找到子节点的最大节点来接任
                val nextValue = maxTreeNode.v.last
//                println("将 ${v}[$childIndex] ${v[childIndex]} -> ${maxTreeNode.v} $nextValue ")
                this.v[childIndex] = nextValue //将当前值设置位接任节点的值

                maxTreeNode.transform()
                if (maxTreeNode.v.isEmpty()) {
                    System.currentTimeMillis() //冷静一下
//                    println("再删除 maxTreeNode.p ${maxTreeNode.p!!.v} 中的 $nextValue")
                    val deleted = maxTreeNode.p!!.v.remove(nextValue)
                    deleted
                } else {
//                    println("再删除 maxTreeNode ${maxTreeNode.v} 中的 $nextValue ")
                    val deleted = maxTreeNode.v.remove(nextValue)
                    deleted
                }
            }
        } else {
            if (this.c.isEmpty()) {
                throw RuntimeException("NOT FOUND $value")
            } else {
                this.c[childIndex].remove(value)
            }
        }

    }

    fun transform() {
        if (this.c.isEmpty() && this.v.isEmpty()) { //等死吧，大兄弟，你已经被抛弃了
            return
        }
        if (this.is2Tree()) { //如果是2-Tree
            if (p != null) { //爸爸还是在的
                val pos = p!!.c.indexOf(this) //找准自己的定位
                val isBigBro: Boolean
                val broIndex = if (pos == p!!.c.size - 1) {
                    isBigBro = false
                    pos - 1 //左侧的兄弟
                } else {
                    isBigBro = true
                    pos + 1 //右侧的兄弟
                }
                val bro = p!!.c[broIndex] //兄弟来了
                if (bro.is2Tree()) { //兄弟是2 Tree
                    if (p!!.is2Tree()) { //爸爸也是个2 tree
                        //todo 有点麻烦，先把爸爸折腾好再来
                        p!!.transform()
                        this.transform()
                    } else {
                        //生成一个4-Tree，yes，4-Tree
                        val bv: Int
                        if (isBigBro) { //需要兄弟忍痛割爱
                            bv = bro.v.pollFirst()
                        } else {
                            bv = bro.v.pollLast()
                        }
                        if (isBigBro) {
                            val pv = p!!.v.removeAt(pos) //取出父亲的值
                            this.v.addLast(pv) //把原来的父亲值合并
                            this.v.addLast(bv) //放入兄弟的值
                            for (tree in bro.c) {
                                this.addLastChild(tree) //放入兄弟的键
                            }
                        } else {
                            val pv = p!!.v.removeAt(broIndex) //取出父亲的值
                            this.v.addFirst(pv) //把原来的父亲值合并
                            this.v.addFirst(bv) //放入兄弟的值
                            bro.c.reverse() //注意顺序
                            for (tree in bro.c) {
                                this.addFirstChild(tree) //放入兄弟的键
                            }
                        }
                        //这个时候，兄弟已经没用了
                        p!!.c.removeAt(broIndex) //割了吧
                    }
                } else {//兄弟是3 Tree
                    val bv: Int
                    val bc: Tree?
                    if (isBigBro) { //需要兄弟忍痛割爱
                        bv = bro.v.pollFirst()
                        bc = bro.c.pollFirst()
                    } else {
                        bv = bro.v.pollLast()
                        bc = bro.c.pollLast()
                    }
                    if (isBigBro) {
                        val pv = p!!.v.set(pos, bv) //取出父亲的值，放入兄弟的值
                        this.v.addLast(pv) //把原来的父亲值合并
                        if (bc != null) {
                            this.addLastChild(bc) //放入兄弟的键
                        }
                    } else {
                        val pv = p!!.v.set(broIndex, bv) //取出父亲的值，放入兄弟的值
                        this.v.addFirst(pv) //把原来的父亲值合并
                        if (bc != null) {
                            this.addFirstChild(bc) //放入兄弟的键
                        }
                    }
                }
            } else {//todo 已经是root
                //已经没得调整，要杀要剐随便
                //完全二叉树，这TM简直了
                //得想个办法
                //好像想到了,儿子什么的，扔掉就好了
                val lc = this.c.pollFirst()
                val rc = this.c.pollFirst()
                if (lc != null && rc != null) {
                    this.v.addFirst(lc.v.pollFirst())
                    this.v.addLast(rc.v.pollFirst())
                    while (lc.c.isNotEmpty()) {
                        this.addChild(lc.c.pollFirst())
                    }
                    while (rc.c.isNotEmpty()) {
                        this.addChild(rc.c.pollFirst())
                    }
                }
            }
        }
    }

    fun delMinLeafToRoot(): Int? {
        val minTreeNode = this.minTreeNode()
        minTreeNode.transform()
        if (minTreeNode.v.isEmpty()) {
            if (minTreeNode.p != null) {
                return minTreeNode.p!!.delMinLeafToRoot()
            } else {
                return null
            }
        }
        return minTreeNode.v.pollFirst()
    }

    fun delMaxLeafToRoot(): Int? {
        val maxTreeNode = this.maxTreeNode()
        maxTreeNode.transform()
        if (maxTreeNode.v.isEmpty()) {
            if (maxTreeNode.p != null) {
                return maxTreeNode.p!!.delMaxLeafToRoot()
            } else {
                println("fuckyou")
            }
        }
        return maxTreeNode.v.pollLast()
    }


    fun delMinRootToLeaf(): Int? {
        return if (this.c.isNotEmpty()) {
            if (this.c[0].is2Tree()) { //左节点是2-Tree
                if (this.c[1].is2Tree()) { //兄弟节点也是2-Tree
                    if (this.is2Tree()) { //自己本身也是2-Tree，（基本就是根节点
                        val l = this.c.pollFirst() //取出左
                        val r = this.c.pollLast() //取出右
                        this.v.addFirst(l.v[0])
                        this.v.add(r.v[0])
                        for (gTree in l.c) {
                            this.addChild(gTree)
                        }
                        for (gTree in r.c) {
                            this.addChild(gTree)
                        }
                        this.delMinRootToLeaf()
                    } else {
                        this.c[0].v.add(this.v.pollFirst())
                        this.c[0].v.add(this.c[1].v.pollFirst())
                        for (gTree in this.c[1].c) {
                            this.c[0].addChild(gTree)
                        }
                        this.c.removeAt(1)
                        this.delMinRootToLeaf()
                    }
                } else if (this.c[1].is3Tree()) {
                    val nv = this.c[1].v.pollFirst() //获取兄弟节点左侧的值
                    val nc: Tree? = this.c[1].c.pollFirst()  //获取兄弟节点左侧的键
                    this.c[0].v.add(this.v.pollFirst()) //把本节点的左值放入左子节点中，使其成为一个2-x节点
                    if (nc != null) { //如果兄弟节点存在键
                        this.c[0].addChild(nc)//那就给其加上
                    }
                    this.v.addFirst(nv) //加上，旋转完成
                    this.c[0].delMinRootToLeaf()
                } else {
                    throw RuntimeException("状态不正确")
                }
            } else { //左节点不是2-Tree
                this.c[0].delMinRootToLeaf()
            }
        } else {
            val pollFirst = this.v.pollFirst()
            format()
            pollFirst
        }
    }

    //todo
    fun format() {
        if (this.c.isEmpty()) { //没有子节点，说明这是底部
            this.p?.format()
        } else {
            if (this.v.size == 3) {
                val l = Tree(this.v.pollFirst(), this.c.pollFirst(), this.c.pollFirst(), this)
                val r = Tree(this.v.pollLast(), this.c.pollFirst(), this.c.pollFirst(), this)
                this.addChild(l)
                this.addChild(r)
                this.p?.upgrade(this)
            } else {
                this.p?.format()
            }
        }
    }


    fun maxTreeNode(): Tree {
        return if (this.c.isEmpty()) {
            this
        } else {
            this.c.last.maxTreeNode()
        }
    }

    fun minTreeNode(): Tree {
        return if (this.c.isEmpty()) {
            this
        } else {
            this.c[0].minTreeNode()
        }
    }


    fun values(): List<Int> {
        return if (this.c.isEmpty()) {
            this.v.toList()
        } else {
            val list = LinkedList<Int>()
            for (i in 0 until this.v.size) {
                list.addAll(c[i].values())
                list.add(v[i])
            }
            list.addAll(this.c.last.values())
            return list
        }

    }

    override fun compareTo(other: Tree): Int {
        val c1 = this.v.first.compareTo(other.v.first)
        if (c1 != 0) {
            return c1
        }
        val c2 = this.v.last.compareTo(other.v.last)
        if (c2 != 0) {
            return c2
        }
        val c3 = this.minTreeNode().v.first.compareTo(other.maxTreeNode().v.last)
        if (c3 != 0) {
            return c3
        }
        val c4 = this.maxTreeNode().v.last.compareTo(other.minTreeNode().v.first)
        if (c4 != 0) {
            return c4
        }
        return 0
    }

    override fun toString(): String {
        return if (this.c.isEmpty()) {
            this.v.joinToString(prefix = "[", postfix = "]", separator = "") { "($it)" }
        } else {
            val sb = StringBuffer("[")
            for (i in 0 until this.v.size) {
                sb.append("${c[i]}(${v[i]})")
            }
            sb.append(this.c.last).append("]")
            sb.toString()
        }
    }

    fun isOK(min: Int = Int.MIN_VALUE, max: Int = Int.MAX_VALUE): Boolean {
        if (this.v.isEmpty()) {
            return true
        }
        return if (this.v.first >= min && this.v.last <= max) {
            if (this.c.isNotEmpty()) {
                val toList = this.v.toList()
                val p = listOf(min) + toList + listOf(max)
                for (i in 0..this.v.size) {
                    if (this.c[i].isOK(p[i], p[i + 1])) {

                    } else {
                        return false
                    }
                }
                return true
            } else {
                true
            }
        } else {
            false
        }

    }


    companion object {
        @JvmStatic
        fun main(args: Array<String>) {
//            val test = "[[[[[[(26)](26)[(27)]](27)[[(28)(28)](28)[(28)]](28)[[(29)](29)[(30)](30)[(30)]]](30)[[[(31)(31)](31)[(32)(32)](32)[(32)]](32)[[(33)](33)[(33)(34)]]](34)[[[(34)](35)[(35)(36)]](36)[[(36)(36)](36)[(37)(37)]]]](37)[[[[(38)(38)](39)[(39)]](39)[[(39)](40)[(40)]]](40)[[[(40)](40)[(40)]](40)[[(41)](41)[(42)(42)]]](42)[[[(43)](43)[(44)]](44)[[(45)](45)[(45)]]]]](46)[[[[[(47)](47)[(47)(47)]](48)[[(49)](49)[(49)](49)[(49)]](50)[[(51)(51)](51)[(51)]]](51)[[[(52)(52)](53)[(53)(53)](54)[(54)]](54)[[(54)](54)[(54)]]]](54)[[[[(55)(55)](55)[(56)]](56)[[(56)](57)[(57)]]](57)[[[(58)](58)[(58)]](58)[[(59)(59)](59)[(59)]]](59)[[[(59)](59)[(60)]](61)[[(61)](61)[(61)]](61)[[(61)](61)[(62)](62)[(62)]]]](63)[[[[(63)](63)[(64)(64)]](64)[[(65)](65)[(65)]]](65)[[[(66)](66)[(66)(67)]](67)[[(68)](69)[(69)(70)]]]]]]"
//            val tree = Tree(test)
//
//            tree.add(53)
//            println(tree.isOK())


            val tree1 = Tree()
            repeat(3) {
                tree1.add(0)
            }
            repeat(4){
                tree1.delMinLeafToRoot()
            }
        }
    }
}