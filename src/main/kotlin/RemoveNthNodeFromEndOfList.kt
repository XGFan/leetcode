fun removeNthFromEnd(head: ListNode?, n: Int): ListNode? {
    var i = n
    var p = head
    val listNode = ListNode(0)
    listNode.next = head
    var t = listNode
    while (p != null) {
        p = p.next
        if (i > 0) {
            i--
        } else {
            t = t.next!!
        }
    }
    t.next = t.next?.next
    return listNode.next
}

fun main() {
    val listNode = ListNode(1)
    var p = listNode
    p.next = ListNode(2)
    p = p.next!!
    p.next = ListNode(3)
    p = p.next!!
    p.next = ListNode(4)
    p = p.next!!
    p.next = ListNode(5)
    p = p.next!!
    println(removeNthFromEnd(listNode, 2))
}
