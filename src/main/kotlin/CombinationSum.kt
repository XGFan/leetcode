fun combinationSum(candidates: IntArray, target: Int): List<List<Int>> {
    combinationSum(candidates, target, emptyList())
    return lists
}

val lists: ArrayList<List<Int>> = ArrayList()

fun combinationSum(
    candidates: IntArray,
    target: Int,
    list: List<Int>,
    offset: Int = 0
) {
    for (i in offset until candidates.size) {
        if (candidates[i] < target) {
            combinationSum(candidates, target - candidates[i], list + candidates[i], i)
        } else if (candidates[i] == target) {
            lists.add(list + candidates[i])
        }
    }
}
