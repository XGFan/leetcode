fun generateParenthesis(n: Int): List<String> {
    val sbs = ArrayList<String>()
    generateParentheses(sbs, "", 0, 0, n)
    return sbs
}

fun generateParentheses(ans: ArrayList<String>, sb: String, left: Int, right: Int, n: Int) {
//    for (i in left + right until 2 * n + 1) {
    if (left == right) {
        if (left < n) {
            generateParentheses(ans, "$sb(", left + 1, right, n)
        } else if (left == n) {
            ans.add(sb)
        }
    } else if (left > right) {
        if (left < n) {
            generateParentheses(ans, "$sb(", left + 1, right, n)
        }
        generateParentheses(ans, "$sb)", left, right + 1, n)
    }
//    }
}

fun main() {
    println(generateParenthesis(3))
}

