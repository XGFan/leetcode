fun countAndSay(n: Int): String {
    var prev = arrayListOf('1')
    for (i in 2..n) {
        val array = ArrayList<Char>(prev.size * 2)
        var c = prev[0]
        var count = 1
        for (j in 1 until prev.size) {
            if (prev[j] != c) {
                array.add((count + 48).toChar())
                array.add(c)
                c = prev[j]
                count = 1
            } else {
                count++
            }
        }
        array.add((count + 48).toChar())
        array.add(c)
        prev = array
    }
    return String(prev.toCharArray())
}


fun main() {
    println(countAndSay(1))
    println(countAndSay(2))
    println(countAndSay(3))
    println(countAndSay(4))
}