fun isMatch2(s: String, p: String): Boolean {
    val t = p.map {
        when (s) {
            "*" -> {
                Wild()
            }
            "?" ->
                Single()
            else -> {
                Normal(s.mapIndexed { index, c -> if (c == it) index else -1 }.filter { it != -1 })
            }
        }
    }
    var p = 0
    var i = 0
    while (true) {
        t[i] is Single
    }
    return false
}

sealed class I {

}

class Normal(val target: List<Int>) : I()
data class Pattern(val min: Int = 0, val max: Int = min)
class Single() : I() {

}

class Wild() : I()