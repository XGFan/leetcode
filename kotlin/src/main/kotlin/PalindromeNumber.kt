fun isPalindrome(x: Int): Boolean {
    if (x == 0) {
        return true
    }
    if (x < 0 || x % 10 == 0) {
        return false
    }
    var x = x
    var r = 0
    while (r < x) {
        val mod = x % 10
        x /= 10
        if (x == r) {
            return true
        }
        r *= 10
        r += mod
        if (x == r) {
            return true
        }
    }
    return false
}

fun main() {
    for (i in 0 until Int.MAX_VALUE) {
        isPalindrome(i)
    }
}