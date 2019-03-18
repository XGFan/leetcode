fun longestPalindrome(s: String): String {
    if (s.isEmpty()) {
        return ""
    }
    var start = 0
    var end = 0
    for (i in 0 until s.length) {
        val len1 = findPalindrome(s, i, i)
        val len2 = findPalindrome(s, i, i + 1)
        val len = Math.max(len1, len2)
        if (len > end - start) {
            start = i - (len - 1) / 2
            end = i + len / 2;
        }
    }
    return s.substring(start, end + 1)
}

fun findPalindrome(s: String, lp: Int, rp: Int): Int {
    var l = lp
    var r = rp
    while (l >= 0 && r < s.length && s[l] == s[r]) {
        l--
        r++
    }
    return r - l - 1
}

fun main() {
    println(longestPalindrome("bb"))
}