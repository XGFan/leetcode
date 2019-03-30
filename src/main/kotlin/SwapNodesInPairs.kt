fun swapPairs(head: ListNode?): ListNode? {
    var p = head
    val answer: ListNode = head?.next ?: return head
    var prev: ListNode? = null
    while (p != null) {
        val temp = p.next
        if (temp != null) {
            p.next = temp.next
            temp.next = p
            if (prev != null) {
                prev.next = temp
            }
        }
        prev = p
        p = p.next
    }
    return answer
}

fun main() {
    println(swapPairs(arrayOf(1, 2, 3, 4).toListNode()))
    println(swapPairs(arrayOf(1, 2, 3).toListNode()))
    println(swapPairs(arrayOf(1, 2).toListNode()))
    println(swapPairs(arrayOf(1).toListNode()))
    println(swapPairs(emptyArray<Int>().toListNode()))
}