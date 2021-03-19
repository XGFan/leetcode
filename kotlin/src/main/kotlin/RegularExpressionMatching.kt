fun isMatch(s: String, p: String): Boolean {
    if (s == p) {
        return true
    }
    var f: Char? = null
    val pList = ArrayList<MatchingChar>()
    for (c in p) {
        if (c == f) {
            pList.last().min++
        } else {
            if (c == '*') {
                pList.last().strict = false
                pList.last().min--
            } else {
                if (c == '.') {
                    pList.add(MatchingChar(c))
                    f = null
                } else {
                    pList.add(MatchingChar(c))
                    f = c
                }
            }
        }
    }
    return match(pList, 0, s, 0)
}

data class MatchingChar(val s: Char, var min: Int = 1, var strict: Boolean = true) {
    var matched = 0

    /**
     * 简单的t/f无法满足需求
     */
    fun match(mc: Char): MatchResult {
        return if (this.strict) {
            if (this.s == mc || this.s == '.') {
                matched++
                when {
                    matched < this.min -> MatchResult.RemoveOther
                    matched == this.min -> MatchResult.RemoveBoth
                    else -> throw RuntimeException("what the fuck $this $mc $matched")
                }
            } else {
                MatchResult.Die
            }
        } else { //动态模式
            if (this.s == mc) { //如果相等
                matched++
                if (matched > this.min) {
                    MatchResult.RemoveItselfOrOtherOrBoth
                } else if (matched == this.min) {
                    MatchResult.RemoveOtherOrBoth
                } else {
                    MatchResult.RemoveOther
                }
            } else if (this.s == '.') {
                MatchResult.RemoveItselfOrOtherOrBoth
            } else {
                if (matched < this.min) {
                    MatchResult.Die
                } else {
                    MatchResult.RemoveItself
                }
            }
        }
    }
}


enum class MatchResult {
    RemoveOther,
    RemoveBoth,
    RemoveItself,
    RemoveOtherOrBoth,
    RemoveItselfOrOtherOrBoth,
    Die
}

fun List<MatchingChar>.deepCopy(): List<MatchingChar> {
    return this.map { it.copy() }.toList()
}

fun matchBlank(p: List<MatchingChar>, pi: Int): Boolean {
    for (index in pi until p.size) {
        if (p[index].strict) {
            return false
        } else {
            if (p[index].min > p[index].matched) {
                return false
            }
        }
    }
    return true

}


fun match(p: List<MatchingChar>, pi: Int, s: String, si: Int): Boolean {
    if (pi >= p.size) {
        return false
    }
    if (si >= s.length) {
        return matchBlank(p, pi)
    }
    return if (pi == p.size - 1 && si == s.length - 1) {
        when (p[pi].match(s[si])) {
            MatchResult.RemoveOther -> false
            MatchResult.RemoveBoth -> true
            MatchResult.RemoveItself -> false
            MatchResult.RemoveOtherOrBoth -> true
            MatchResult.RemoveItselfOrOtherOrBoth -> true
            MatchResult.Die -> false
        }
    } else {
        when (p[pi].match(s[si])) {
            MatchResult.RemoveOther -> match(p, pi, s, si + 1)
            MatchResult.RemoveBoth -> match(p, pi + 1, s, si + 1)
            MatchResult.RemoveItself -> match(p, pi + 1, s, si)
            MatchResult.RemoveOtherOrBoth -> match(p.deepCopy(), pi, s, si + 1)
                    || match(p.deepCopy(), pi + 1, s, si + 1)
            MatchResult.RemoveItselfOrOtherOrBoth -> match(p.deepCopy(), pi, s, si + 1)
                    || match(p.deepCopy(), pi + 1, s, si)
                    || match(p.deepCopy(), pi + 1, s, si + 1)
            MatchResult.Die -> false
        }
    }
}

fun main() {
    println(isMatch("aa", "a")) //false
    println(isMatch("aa", "a*")) //true
    println(isMatch("ab", ".*"))//true
    println(isMatch("aab", "c*a*b"))//true
    println(isMatch("mississippi", "mis*is*p*."))//false
    println(isMatch("aaa", "a.a"))//true
    println(isMatch("aaa", "a*a"))//true
    println(isMatch("aaa", "ab*a*c*a"))//true
    println(isMatch("a", "ab*"))//true
    println(isMatch("a", ".*..a*"))//false
    println(isMatch("ab", ".*.."))//true
    println(isMatch("", ""))//true
    println(isMatch("bbbaccbbbaababbac", ".b*b*.*...*.*c*."))//true

}