fun isValidSudoku(board: Array<CharArray>): Boolean {
    val map = Array(27) { IntArray(9) }
    for (i in 0 until 9) {
        for (j in 0 until 9) {
            val c = board[i][j]
            if (c != '.') {
                val toInt = c.toInt() - 49
                if (++map[i][toInt] > 1) {
                    return false
                }
                if (++map[9 + j][toInt] > 1) {
                    return false
                }
                if (++map[((j / 3) * 3) + (i / 3) + 18][toInt] > 1) {
                    return false
                }
            }
        }
    }
    return true
}