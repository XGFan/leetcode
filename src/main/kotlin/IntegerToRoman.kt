fun intToRoman(num: Int): String {
    val sb = StringBuilder()
    when (num / 1000) {
        1 -> sb.append("M")
        2 -> sb.append("MM")
        3 -> sb.append("MMM")
    }
    when (num / 100 % 10) {
        1 -> sb.append("C")
        2 -> sb.append("CC")
        3 -> sb.append("CCC")
        4 -> sb.append("CD")
        5 -> sb.append("D")
        6 -> sb.append("DC")
        7 -> sb.append("DCC")
        8 -> sb.append("DCCC")
        9 -> sb.append("CM")
    }
    when (num / 10 % 10) {
        1 -> sb.append("X")
        2 -> sb.append("XX")
        3 -> sb.append("XXX")
        4 -> sb.append("XL")
        5 -> sb.append("L")
        6 -> sb.append("LX")
        7 -> sb.append("LXX")
        8 -> sb.append("LXXX")
        9 -> sb.append("XC")
    }
    when (num % 10) {
        1 -> sb.append("I")
        2 -> sb.append("II")
        3 -> sb.append("III")
        4 -> sb.append("IV")
        5 -> sb.append("V")
        6 -> sb.append("VI")
        7 -> sb.append("VII")
        8 -> sb.append("VIII")
        9 -> sb.append("IX")
    }
    return sb.toString()
}

fun setValue(sb: StringBuilder, num: Int, i: String, v: String, x: String) {
    when {
        num < 4 -> {
            repeat(num) {
                sb.append(i)
            }
        }
        num == 4 -> {
            sb.append(i)
            sb.append(v)
        }
        num in 5..8 -> {
            sb.append(v)
            repeat(num - 5) {
                sb.append(i)
            }
        }
        num == 9 -> {
            sb.append(i)
            sb.append(x)
        }
    }
}