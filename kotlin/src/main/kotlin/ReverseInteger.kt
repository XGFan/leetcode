fun reverse(x: Int): Int {
    var x = x
    var r = 0
    while (x != 0) {
        if (r != r * 10 / 10) {
            return 0
        }
        r *= 10
        r += x % 10
        x /= 10
    }
    return r
}

fun main() {
    println(reverse(100))
    println(reverse(-201))
    println(reverse(-320))
    println(reverse(1534236469))
}