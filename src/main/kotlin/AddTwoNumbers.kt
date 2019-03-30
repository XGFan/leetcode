fun addTwoNumbers(l1: ListNode?, l2: ListNode?): ListNode? {
    var p = l1
    var q = l2
    var up = 0
    var last: ListNode? = null
    while (p != null || q != null) {
        var sum = (p?.`val` ?: 0) + (q?.`val` ?: 0) + up
        up = 0
        if (sum >= 10) {
            sum -= 10
            up = 1
        }
        val who = p ?: q
        who!!.`val` = sum
        if (last == null) {
            last = who
        } else {
            last.next = who
            last = last.next
        }
        p = p?.next
        q = q?.next
    }
    if (up == 1) {
        last?.next = ListNode(1)
    }
    return l1
}


/**
 * 理解错题目了
 */
fun addTwoNumbersStpuid(l1: ListNode?, l2: ListNode?): ListNode? {
    var rl1: ListNode? = reverseListNode(l1)
    var rl2: ListNode? = reverseListNode(l2)
    var q: ListNode? = null
    var up = 0
    while (rl1 != null || rl2 != null) {
        var sum = (rl2?.`val` ?: 0) + (rl1?.`val` ?: 0) + up
        up = 0
        if (sum >= 10) {
            sum -= 10
            up = 1
        }
        var temp: ListNode
        if (rl1 != null) {
            rl1.`val` = sum
            temp = rl1
        } else {
            rl2?.`val` = sum //其实不用相加了
            temp = rl2!!
        }
        rl2 = rl2?.next
        rl1 = rl1?.next
        temp.next = q
        q = temp
    }
    return q
}

fun reverseListNode(list: ListNode?): ListNode? {
    var p: ListNode? = list
    var q: ListNode? = null
    while (p != null) {
        val temp = p
        p = p.next
        temp.next = q
        q = temp
    }
    return q
}

class ListNode(var `val`: Int) {
    var next: ListNode? = null

    override fun toString(): String {
        return if (next == null) {
            "$`val`"
        } else {
            "$`val`,$next"
        }
    }

}


fun Array<Int>.toListNode(): ListNode? {
    if (this.isEmpty()) {
        return null
    }
    val a = ListNode(this[0])
    this.drop(1).map { ListNode(it) }.fold(a) { a, b ->
        a.next = b
        a.next!!
    }
    return a
}

fun main() {
    var p = ListNode(0)
    val listNode = p
    for (i in 1..9) {
        p.next = ListNode(i)
        p = p.next!!
    }
    println(listNode)
    println(reverseListNode(listNode))
    println(listNode)

}