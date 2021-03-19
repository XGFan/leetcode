fun mergeKLists(lists: Array<ListNode?>): ListNode? {
    val head: ListNode = ListNode(0)
    var p = head
    val flag = ListNode(Int.MAX_VALUE)
    while (true) {
        var value: Int? = null
        var index = -1
        for (j in 0 until lists.size) {
            if (lists[j] != flag && lists[j] != null && (value == null || value > lists[j]!!.`val`)) {
                value = lists[j]!!.`val`
                index = j
            }
        }
        if (index == -1) {
            break
        }
        p.next = lists[index]!!
        p = p.next!!
        lists[index] = lists[index]!!.next ?: flag
    }
    return head.next
}

fun main() {
    //[[1,4,5],[1,3,4],[2,6]]
    println(
        mergeKLists(
            arrayOf(null)
        )
    )
    println(
        mergeKLists(
            arrayOf(
                arrayOf(1, 4, 5).toListNode(),
                arrayOf(1, 3, 4).toListNode(),
                arrayOf(2, 6).toListNode()
            )
        )
    )
}