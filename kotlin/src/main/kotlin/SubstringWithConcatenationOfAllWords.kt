fun findSubstring(s: String, words: Array<String>): List<Int> {
    if (words.isEmpty())
        return emptyList()

    val counter = HashMap<String, Int>()
    for (word in words) {
        counter[word] = (counter[word] ?: 0) + 1
    }
    val wordLen = words[0].length
    val arrays = Array<String?>(s.length) { null }
    for (i in 0 until s.length - wordLen + 1) {
        word@ for (word in counter.keys) {
            for (k in 0 until wordLen) {
                if (word[k] != s[i + k]) {
                    continue@word
                }
            }
            arrays[i] = word
            break
        }
    }
    val arrayList = ArrayList<Int>()
    val flag = HashMap<String, Int>()
    loop@ for (index in s.indices) {
        flag.clear()
        if (arrays[index] == null) {
            continue
        }
        for (offset in words.indices) {
            val word = arrays.getOrNull(index + offset * wordLen)
            if (word != null) {
                flag[word] = (flag[word] ?: 0) + 1
            } else {
                continue@loop
            }
        }
        if (flag == counter) {
            arrayList.add(index)
        }
    }
    return arrayList
}

fun main() {
    println(findSubstring("barfoothefoobar", arrayOf("foo", "bar")))
    println(findSubstring("wordgoodgoodgoodbestword", arrayOf("word", "good", "best", "word")))
    println(findSubstring("wordgoodgoodgoodbestword", arrayOf("word", "good", "best", "good")))
//    System.exit(0)
    println(
        findSubstring(
            String(List(5000) { 'a' }.toCharArray()),
            List(5000) { "a" }.toTypedArray()
        )
    )
}