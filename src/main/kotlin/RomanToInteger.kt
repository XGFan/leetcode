fun romanToInt(s: String): Int {
    var sum = 0
    var last = -1
    for (c in s) {
        val i = map[c]!!
        sum += if (i / last == 5 || i / last == 10) {
            i - last - last
        } else {
            i
        }
        last = i
    }
    return sum
}

val map = mapOf(
    'I' to 1,
    'V' to 5,
    'X' to 10,
    'L' to 50,
    'C' to 100,
    'D' to 500,
    'M' to 1000
)