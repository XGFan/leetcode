import java.util.*

object T {
    fun randomTestAll() {
        val list = LinkedList(listOf(50))
        val tree = Tree(50)
        var addOp: Long = 0
        var delOp: Long = 0
        var mOp: Long = 0
        var i: Long = 0

        val random = Random()
        loop@ while (true) {
            i++
            val op = random.nextInt(60)
            val pinnedTree = tree.toString()
            when (op) {
                in 0..25 -> {
                    val value = list.pollFirst() ?: continue@loop
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
                in 25..30 -> { //getAndRemoveMin
                    val min = list.min()
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
                    if (list.size >= 10_000 || i >= 1_000_000) {
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
                        this.p?.upgrade(this)
                    }
                    "1-2", "2-3", "3-4" -> this.c[getIndexInChild(tree.v[0])].addTree(tree)  //full-tree存在子节点，让子节点去处理
                    else -> throw RuntimeException("状态不正确，添加${tree.type}到${this.type}")
                }
            }
            "1-2" -> {//添加一个full-2-tree
                this.v.add(childIndex, tree.v[0])
                for (t in tree.c.reversed()) {
                    this.addChild(childIndex, t)
                }
                when (this.type) {
                    "2-3" -> {//full-2-tree（某个子节点进位) -> full-3-tree -> keep
                    }
                    "3-4" -> {//full-3-tree（某个子节点进位) -> full-4-tree -> 2 full-2-tree
                        val l = Tree(this.v.pollFirst(), this.c.pollFirst(), this.c.pollFirst(), this)
                        val r = Tree(this.v.pollLast(), this.c.pollFirst(), this.c.pollFirst(), this)
                        this.addChild(l)
                        this.addChild(r)
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
                    this.v.remove(value)
                }
            } else {
                val maxTreeNode = this.c[childIndex].maxTreeNode() //找到子节点的最大节点来接任
                val nextValue = maxTreeNode.v.last
                this.v[childIndex] = nextValue //将当前值设置位接任节点的值
                maxTreeNode.transform()
                if (maxTreeNode.v.isEmpty()) {
                    maxTreeNode.p!!.v.remove(nextValue)
                } else {
                    maxTreeNode.v.remove(nextValue)
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
            return minTreeNode.p?.delMinLeafToRoot()
        }
        return minTreeNode.v.pollFirst()
    }

    fun delMaxLeafToRoot(): Int? {
        val maxTreeNode = this.maxTreeNode()
        maxTreeNode.transform()
        if (maxTreeNode.v.isEmpty()) {
            return maxTreeNode.p!!.delMaxLeafToRoot()
        }
        return maxTreeNode.v.pollLast()
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
}