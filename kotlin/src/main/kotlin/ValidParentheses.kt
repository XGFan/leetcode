fun isValid(s: String): Boolean {
    val arrayList = ArrayList<Char>()
    for (c in s) {
        val t = arrayList.getOrNull(arrayList.lastIndex)
        if (t != null) {
            when (c) {
                '{', '[', '(' -> arrayList.add(c)
                '}' -> if (t == '{') {
                    arrayList.removeAt(arrayList.lastIndex)
                } else {
                    return false
                }
                ']' -> if (t == '[') {
                    arrayList.removeAt(arrayList.lastIndex)
                } else {
                    return false
                }
                ')' -> if (t == '(') {
                    arrayList.removeAt(arrayList.lastIndex)
                } else {
                    return false
                }
                else -> return false
            }
        } else {
            arrayList.add(c)
        }
    }
    return arrayList.isEmpty()
}

fun main() {
    println(isValid("[{}]"))
    println(isValid("[}]"))
    println(isValid("[]"))
    println(isValid("["))
    println(isValid(""))
}