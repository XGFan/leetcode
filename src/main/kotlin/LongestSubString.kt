fun lengthOfLongestSubstring0(s: String): Int {
    val map = HashMap<Char, Int>()
    val intArray = IntArray(s.length)
    var lock = -1
    s.forEachIndexed { index, c ->
        val prev = map[c] //找到之前的下标
        if (prev != null) {
            lock = maxOf(prev, lock)
        }
        for (i in lock + 1..index) {
            intArray[i]++
        }
        map[c] = index
    }
    return intArray.max() ?: 0
}

fun lengthOfLongestSubstring(s: String): Int {
    val chars = HashMap<Char, Int>()
    var lock = -1
    var maxLen = 0
    s.forEachIndexed { index, char ->
        val last = chars[char] //上一次出现的位置
        if (last != null) {
            lock = maxOf(lock, last) //使用新🔒
        }
        maxLen = maxOf(maxLen, index - lock) //比较两个相同节点之间距离与最大距离
        chars[char] = index
    }
    return maxLen
}

fun main() {
    println(lengthOfLongestSubstring("tmmzuxt"))
    println(lengthOfLongestSubstring0("tmmzuxt"))
    println(lengthOfLongestSubstring("museuwzbczdejn"))
    println(lengthOfLongestSubstring0("museuwzbczdejn"))
//    museuwzbczdejn
//    43768767654321
}