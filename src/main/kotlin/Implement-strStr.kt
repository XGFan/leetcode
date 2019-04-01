fun strStr(haystack: String, needle: String): Int {
    loop@ for (i in 0 until (haystack.length - needle.length) + 1) {
        for (j in needle.indices) {
            if (haystack[i + j] != needle[j]) {
                continue@loop
            }
        }
        return i
    }
    return -1
}

fun main() {
    println(strStr("aaa", "aaaa"))
    println(strStr("aaa", ""))
}