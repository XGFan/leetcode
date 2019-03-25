val phoneChar = mapOf(
    '2' to arrayOf('a', 'b', 'c'),
    '3' to arrayOf('d', 'e', 'f'),
    '4' to arrayOf('g', 'h', 'i'),
    '5' to arrayOf('j', 'k', 'l'),
    '6' to arrayOf('m', 'n', 'o'),
    '7' to arrayOf('p', 'q', 'r', 's'),
    '8' to arrayOf('t', 'u', 'v'),
    '9' to arrayOf('w', 'x', 'y', 'z')
)

fun letterCombinations(digits: String): List<String> {
    if (digits.isEmpty()) {
        return emptyList()
    }
    return digits.map { phoneChar[it] ?: emptyArray() }.fold(listOf("")) { acc: List<String>, chars: Array<Char> ->
        acc.flatMap { s ->
            chars.map { s.plus(it) }
        }
    }
}