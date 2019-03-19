fun myAtoi(str: String): Int {
    var x = 0
    var negative = false
    var canBeIgnore = true
    loop@ for (c in str) {
        val intRange = 48 until 58
        val toInt = c.toInt()
        when (toInt) {
            45 -> if (canBeIgnore) {
                negative = true
                canBeIgnore = false
            } else {
                break@loop
            }
            43 -> if (canBeIgnore) {
                negative = false
                canBeIgnore = false
            } else {
                break@loop
            }
            32 -> if (!canBeIgnore) break@loop
            in intRange -> {
                val t = x
                x *= 10
                x += toInt - 48
                if (t != x / 10) {
                    return if (negative) Int.MIN_VALUE else Int.MAX_VALUE
                }
                canBeIgnore = false
            }
            else -> {
                break@loop
            }
        }
    }
    return if (negative) -x else x
}

fun nullOnFalse(status: Boolean): Any? {
    return if (status) {
        Any()
    } else {
        null
    }
}

fun main() {
    println(myAtoi("42"))
    println(myAtoi("   -42"))
    println(myAtoi("4193 with words"))
    println(myAtoi("words and 987"))
    println(myAtoi("-91283472332"))
    println(myAtoi("2147483648"))
}