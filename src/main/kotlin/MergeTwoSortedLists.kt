fun mergeTwoLists(l1: ListNode?, l2: ListNode?): ListNode? {
    var l1 = l1
    var l2 = l2
    val head: ListNode = ListNode(0)
    var p = head
    while (l1 != null && l2 != null) {
        if (l1.`val` <= l2.`val`) {
            p.next = l1
            l1 = l1.next
        } else {
            p.next = l2
            l2 = l2.next
        }
        p = p.next!!
    }
    if (l1 != null) {
        p.next = l1
        return head.next
    }
    if (l2 != null) {
        p.next = l2
        return head.next
    }
    return head.next
}

fun main() {
    val a = ListNode(1)
    listOf(2, 4).map { ListNode(it) }.fold(a) { a, b ->
        a.next = b
        a.next!!
    }

    val b = ListNode(1)
    listOf(2, 3).map { ListNode(it) }.fold(b) { a, b ->
        a.next = b
        a.next!!
    }
    println(a)
    println(b)

    println(mergeTwoLists(a, b))

}