fun longestCommonPrefix(strs: Array<String>): String {
    if (strs.isEmpty()) {
        return ""
    }
    for (i in 0 until strs[0].length) {
        for (j in 1 until strs.size) {
            if (strs[0][i] != strs[j].getOrNull(i) ?: return strs[0].substring(0, i)) {
                return strs[0].substring(0, i)
            }
        }
    }
    return strs[0]
}

fun main() {
    println(longestCommonPrefix(arrayOf("flower", "flow", "flight")))
    println(longestCommonPrefix(arrayOf()))
}