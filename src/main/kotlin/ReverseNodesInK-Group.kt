fun reverseKGroup(head: ListNode?, k: Int): ListNode? {
    var p = head
    var ans: ListNode? = null
    var t: ListNode? = null
    while (p != null) {
        val newHead = reverseOrNot(p, k) //newHead 为头，p为尾巴
        if (ans == null) { //第一次遍历
            ans = newHead //ans 指向头
        } else {
            t?.next = newHead //连上
        }
        t = p //t指向尾巴
        p = p.next
    }
    return ans
}

fun reverseOrNot(head: ListNode?, size: Int): ListNode? {
    var h = head
    var p = head?.next
    var i = 1
    while (p != null && i < size) {
        head?.next = p.next //head实际上是tail
        p.next = h //当前点p的next指向旧head h
        h = p //新的head h指向p
        p = head?.next //p指向下一个点
        i++
    }
    return if (i < size) {
        reverseOrNot(h, i) //如果长度不够，干脆再翻转一次
    } else {
        h
    }
}

fun main() {
//    println(reverseOrNot(arrayOf(1, 2, 3, 4).toListNode(), 5))
//    println(reverseOrNot(arrayOf(1, 2, 3, 4).toListNode(), 4))
//    println(reverseOrNot(arrayOf(1, 2, 3, 4).toListNode(), 3))
//    println(reverseOrNot(arrayOf(1, 2, 3, 4).toListNode(), 2))
//    println(reverseOrNot(arrayOf(1, 2, 3, 4).toListNode(), 1))
//    println(reverseOrNot(arrayOf(1, 2, 3, 4).toListNode(), 0))


    println(reverseKGroup(arrayOf(1, 2, 3, 4, 5).toListNode(), 1))
    println(reverseKGroup(arrayOf(1, 2, 3, 4, 5).toListNode(), 2))
    println(reverseKGroup(arrayOf(1, 2, 3, 4, 5).toListNode(), 3))
    println(reverseKGroup(arrayOf(1, 2, 3, 4, 5).toListNode(), 4))
    println(reverseKGroup(arrayOf(1, 2, 3, 4, 5).toListNode(), 5))
    println(reverseKGroup(arrayOf(1, 2, 3, 4, 5).toListNode(), 6))
}
