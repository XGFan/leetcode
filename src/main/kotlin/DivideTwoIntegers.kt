fun divide(dividend: Int, divisor: Int): Int {
    if (divisor == 1) {
        return dividend
    }
    if (divisor == -1) {
        return if (dividend == Int.MIN_VALUE) {
            Int.MAX_VALUE
        } else {
            -dividend
        }
    }
    if (dividend == divisor) {
        return 1
    }
    var count = 0
    var last = dividend
    return if (dividend > 0) {
        if (divisor > 0) {
            count = -1
            while (last >= 0) {
                last -= divisor
                count++
            }
            count
        } else {
            count = 1
            while (last >= 0) {
                last += divisor
                count--
            }
            count
        }
    } else {
        if (divisor > 0) {
            count = 1
            while (last <= 0) {
                last += divisor
                count--
            }
            count
        } else {
            count = -1
            while (last <= 0) {
                last -= divisor
                count++
            }
            count
        }
    }
}

fun main() {
    println(divide(-2147483648, -2147483648))
    println(divide(-2147483648, -1))
    println(divide(-2147483648, 1))
    println(divide(2147483647, 1))
    println(divide(2147483647, -1))
    println(divide(2147483647, 2))
    println(divide(10, 3))
    println(divide(7, -3))
    println(divide(-9, -3))
    println(divide(-8, -3))
    println(divide(-8, 2))
    println(divide(0, 1))
}