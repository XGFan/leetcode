fun combinationSum2(candidates: IntArray, target: Int): List<List<Int>> {
    combinationSum2(candidates.sortedArray(), target, emptyList())
    return lists2
}

val lists2: ArrayList<List<Int>> = ArrayList()

fun combinationSum2(
    candidates: IntArray,
    target: Int,
    list: List<Int>,
    offset: Int = 0
) {
    var last = 0
    for (i in offset until candidates.size) {
        if (candidates[i] != last) {
            last = candidates[i]
            if (candidates[i] < target) {
                combinationSum2(candidates, target - candidates[i], list + candidates[i], i + 1)
            } else if (candidates[i] == target) {
                lists2.add(list + candidates[i])
            }
        }
    }
}

fun main() {
    val intArrayOf = intArrayOf(10, 1, 2, 7, 6, 1, 5)
    println(combinationSum2(intArrayOf, 8).map { it.toString() }.joinToString("\n"))
}