fun longestValidParentheses(s: String): Int {
    val temp = ArrayList<Pair<Char, Int>>()
    val ranges = ArrayList<Pair<Int, Int>>()
    for (index in s.indices) {
        when (s[index]) {
            '(' -> {
                temp.add(Pair(s[index], index))
            }
            ')' -> {
                if (temp.lastOrNull()?.first == '(') {
                    ranges.diy(Pair(temp.last().second, index))
                    temp.removeAt(temp.lastIndex)
                } else {
                    temp.clear()
                }
            }
            else -> throw RuntimeException("?")
        }
    }
    var sum = 0
    for (range in ranges) {
        sum = Math.max(range.second - range.first + 1, sum)
    }
    return sum
}


fun ArrayList<Pair<Int, Int>>.diy(newOne: Pair<Int, Int>) {
    when {
        newOne.first - 1 == this.lastOrNull()?.second -> { //()()
            this[this.lastIndex] = this[this.lastIndex].copy(second = newOne.second)
        }
        newOne.first + 1 == this.lastOrNull()?.first -> {//(())
            this.removeAt(this.lastIndex)
            this.diy(newOne)
        }
        else -> this.add(newOne)
    }
}

fun main() {
    println(longestValidParentheses("(()"))
    println(longestValidParentheses(")()())"))
    println(longestValidParentheses("()(())"))
    println(longestValidParentheses(""))
    println(longestValidParentheses("("))
}