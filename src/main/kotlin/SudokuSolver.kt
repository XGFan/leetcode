fun solveSudoku(board: Array<CharArray>) {
    val sudoku = Array(9) { IntArray(9) { 511 } }
    for (i in 0 until 9) {
        for (j in 0 until 9) {
            val c = board[i][j]
            if (c != '.') {
                reduce(i, j, c.toInt() - 48, sudoku)
            }
        }
    }

//    if (sudoku.STATUS != 1) {
//        println()
//        val justTry = justTry(sudoku)!!
//        for (i in 0 until 9) {
//            for (j in 0 until 9) {
//                board[i][j] = ((justTry[i][j] - 512).onlyValue() + 48).toChar()
//            }
//        }
//    } else {
//        for (i in 0 until 9) {
//            for (j in 0 until 9) {
//                board[i][j] = ((sudoku[i][j] - 512).onlyValue() + 48).toChar()
//            }
//        }
//    }

    for (i in 0 until 9) {
        for (j in 0 until 9) {
            findOnlyValue(i, j, sudoku)
        }
    }
    val temp = Array(9) { Array<String>(9) { "" } }

    anotherReduce(sudoku)

    for (i in 0 until 9) {
        for (j in 0 until 9) {
            val x = sudoku[i][j]
            temp[i][j] = if ((x - 512).onlyValue() == 0) {
                x.possibles().map { (it + 48).toChar() }.joinToString("|", "[", "]")
            } else {
                ((x - 512).onlyValue() + 48).toChar().toString()
            }
        }
    }
    println(temp.map { it.joinToString(",") }.joinToString("\n"))

}

fun justTry(sudoku: Array<IntArray>): Array<IntArray>? {
    for (i in 0 until 9) {
        for (j in 0 until 9) {
            if (sudoku[i][j] < 512) {
                val maybeArray = sudoku[i][j].possibles() //需要try的部分
                if (maybeArray.size > 1) {
                    for (test in maybeArray) {
                        val copy = sudoku.map { it.copyOf() }.toTypedArray()
                        reduce(i, j, test, copy) //根据此次尝试，过一遍
                        val status = copy.STATUS
                        return if (status == -1) {
                            sudoku[i][j] = sudoku[i][j] and (1 shl (test - 1)).inv()
                            continue//再次尝试
                        } else if (status == 0) {
                            justTry(copy) ?: continue //可能是正确的，接着尝试。如果尝试不出来，就再次尝试
                        } else {
                            copy
                        }
                    }
                }
            }
        }
    }
    return if (sudoku.STATUS == 1) {
        sudoku
    } else {
        null
    }
}

/**
 * 1 : ok
 * 0 : possible
 * -1 : wrong
 */
val Array<IntArray>.STATUS: Int
    get() {
        var status = 1
        for (ints in this) {
            for (int in ints) {
                if (int == 0 || int == 512) {
                    return -1
                } else if (int < 512) {
                    status = 0
                }
            }
        }
        return status
    }

/**
 * 移除各种可能性
 * 十位二进制，第十位表示是否已经固定，
 * 比如:
 * 1000000001 表示为1
 * 1000000010 表示为2
 * 1000000100 表示为3
 * 0111111111 表示可能为1-9
 */
fun reduce(i: Int, j: Int, value: Int, sudoku: Array<IntArray>) {
    sudoku[i][j] = (1 shl (value - 1)) + 512 //第十位表示已经处理过一遍了,向左移动value-1位
    for (x in 0 until 9) {
        val needToBeRemoved = (1 shl (value - 1)).inv()
        //去除横向
        if (x != j) {
            sudoku[i][x] = sudoku[i][x] and needToBeRemoved
            if (sudoku[i][x].onlyValue() != 0 && sudoku[i][x] < 512) {
                reduce(i, x, sudoku[i][x].onlyValue(), sudoku)
            }
        }
        //去除竖向
        if (x != i) {
            sudoku[x][j] = sudoku[x][j] and needToBeRemoved
            if (sudoku[x][j].onlyValue() != 0 && sudoku[x][j] < 512) {
                reduce(x, j, sudoku[x][j].onlyValue(), sudoku)
            }
        }
        //去除九宫格可能性
        val ni = i / 3 * 3 + x / 3
        val nj = j / 3 * 3 + x % 3
        if (ni != i || nj != j) {
            sudoku[ni][nj] = sudoku[ni][nj] and needToBeRemoved
            if (sudoku[ni][nj].onlyValue() != 0 && sudoku[ni][nj] < 512) {
                reduce(ni, nj, sudoku[ni][nj].onlyValue(), sudoku)
            }
        }
    }
}

fun findOnlyValue(i: Int, j: Int, sudoku: Array<IntArray>) {
    val value = sudoku[i][j]
    if (value >= 512) {
        return
    }
    var temp = value
    for (offset in 0 until 9) {
        //去除横向
        if (offset != j) {
            temp = sudoku[i][offset].inv() and temp
        }
    }
    if (temp.possibles().size == 1) {
//        println("I found it! ($i,$j) = ${temp.onlyValue()}")
        reduce(i, j, temp.onlyValue(), sudoku)
        return
    }
    temp = value
    for (offset in 0 until 9) {
        //去除横向
        if (offset != i) {
            temp = sudoku[offset][j].inv() and temp
        }
    }
    if (temp.possibles().size == 1) {
//        println("I found it! ($i,$j) = ${temp.onlyValue()}")
        reduce(i, j, temp.onlyValue(), sudoku)
        return
    }
    temp = value
    for (offset in 0 until 9) {
        //去除九宫格可能性
        val ni = i / 3 * 3 + offset / 3
        val nj = j / 3 * 3 + offset % 3
        if (ni != i || nj != j) {
            temp = sudoku[ni][nj].inv() and temp
        }
    }
    if (temp.possibles().size == 1) {
//        println("I found it! ($i,$j) = ${temp.onlyValue()}")
        reduce(i, j, temp.onlyValue(), sudoku)
        return
    }

}

fun anotherReduce(sudoku: Array<IntArray>) {
    //先横着
    for (x in 0 until 9) {
        for (y in 0 until 3) {
            val a = sudoku[x][y * 3]
            val b = sudoku[x][y * 3 + 1]
            val c = sudoku[x][y * 3 + 2]
            val theSameValue: Int? =
                if (a == b && b == c && a.possibles().size == 3) {
                    a
                } else {
                    if (a == b && a.possibles().size == 2) {
                        a
                    } else if (b == c && b.possibles().size == 2) {
                        b
                    } else if (a == c && c.possibles().size == 2) {
                        c
                    } else {
                        null
                    }
                }
            if (theSameValue != null) {
                for (z in 0 until 9) {
                    //去除横向
                    if (!(z == y * 3 || z == y * 3 + 1 || z == y * 3 + 2)) {
                        sudoku[x][z] = sudoku[x][z] and theSameValue.inv()
                        if (sudoku[x][z].onlyValue() != 0 && sudoku[x][z] < 512) {
                            reduce(x, z, sudoku[x][z].onlyValue(), sudoku)
                        }
                    }
                    //没有办法去除竖向
                    //去除九宫格可能性
                    val ni = x / 3 * 3 + z / 3
                    val nj = y * 3 / 3 * 3 + z % 3
                    if (ni != x) {
                        sudoku[ni][nj] = sudoku[ni][nj] and theSameValue.inv()
                        if (sudoku[ni][nj].onlyValue() != 0 && sudoku[ni][nj] < 512) {
                            reduce(ni, nj, sudoku[ni][nj].onlyValue(), sudoku)
                        }
                    }
                }
            }

        }

    }
    for (y in 0 until 9) {
        for (x in 0 until 3) {
            val a = sudoku[x * 3][y]
            val b = sudoku[x * 3 + 1][y]
            val c = sudoku[x * 3 + 2][y]
            val theSameValue: Int? =
                if (a == b && b == c && a.possibles().size == 3) {
                    a
                } else {
                    if (a == b && a.possibles().size == 2) {
                        a
                    } else if (b == c && b.possibles().size == 2) {
                        b
                    } else if (a == c && c.possibles().size == 2) {
                        c
                    } else {
                        null
                    }
                }
            if (theSameValue != null) {
                for (z in 0 until 9) {
                    //去除横向
                    if (!(z == x * 3 || z == x * 3 + 1 || z == x * 3 + 2)) {
                        sudoku[z][y] = sudoku[z][y] and theSameValue.inv()
                        if (sudoku[z][y].onlyValue() != 0 && sudoku[z][y] < 512) {
                            reduce(z, y, sudoku[z][y].onlyValue(), sudoku)
                        }
                    }
                    //去除九宫格可能性
                    val ni = x * 3 / 3 * 3 + z / 3
                    val nj = y / 3 * 3 + z % 3
                    if (nj != y) {
                        sudoku[ni][nj] = sudoku[ni][nj] and theSameValue.inv()
                        if (sudoku[ni][nj].onlyValue() != 0 && sudoku[ni][nj] < 512) {
                            reduce(ni, nj, sudoku[ni][nj].onlyValue(), sudoku)
                        }
                    }
                }
            }
        }
    }
}

fun Int.onlyValue(): Int {
    return when (this) {
        1 -> 1
        2 -> 2
        4 -> 3
        8 -> 4
        16 -> 5
        32 -> 6
        64 -> 7
        128 -> 8
        256 -> 9
        else -> 0
    }
}

fun Int.possibles(): List<Int> {
    val arrayList = ArrayList<Int>()
    var c = when {
        this >= 512 -> this - 512
        else -> this
    }
    var count = 1
    while (c > 0) {
        val i = c and 1
        if (i == 1) {
            arrayList.add(count)
        }
        c = c shr 1
        count++
    }
    return arrayList
}

fun main() {
    val fucks = listOf(
//        """[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]""",
        """[[".",".","9","7","4","8",".",".","."],["7",".",".",".",".",".",".",".","."],[".","2",".","1",".","9",".",".","."],[".",".","7",".",".",".","2","4","."],[".","6","4",".","1",".","5","9","."],[".","9","8",".",".",".","3",".","."],[".",".",".","8",".","3",".","2","."],[".",".",".",".",".",".",".",".","6"],[".",".",".","2","7","5","9",".","."]]""",
        """[[".",".",".","2",".",".",".","6","3"],["3",".",".",".",".","5","4",".","1"],[".",".","1",".",".","3","9","8","."],[".",".",".",".",".",".",".","9","."],[".",".",".","5","3","8",".",".","."],[".","3",".",".",".",".",".",".","."],[".","2","6","3",".",".","5",".","."],["5",".","3","7",".",".",".",".","8"],["4","7",".",".",".","1",".",".","."]]""",
        """[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]""",
        """[["1",".",".",".","7",".",".","3","."],["8","3",".","6",".",".",".",".","."],[".",".","2","9",".",".","6",".","8"],["6",".",".",".",".","4","9",".","7"],[".","9",".",".",".",".",".","5","."],["3",".","7","5",".",".",".",".","4"],["2",".","3",".",".","9","1",".","."],[".",".",".",".",".","2",".","4","3"],[".","4",".",".","8",".",".",".","9"]]""",
        """[[".",".","9","7","4","8",".",".","."],["7",".",".",".",".",".",".",".","."],[".","2",".","1",".","9",".",".","."],[".",".","7",".",".",".","2","4","."],[".","6","4",".","1",".","5","9","."],[".","9","8",".",".",".","3",".","."],[".",".",".","8",".","3",".","2","."],[".",".",".",".",".",".",".",".","6"],[".",".",".","2","7","5","9",".","."]]"""
    )
    for (fuck in fucks) {
        val map = fuck.drop(1).dropLast(1).split("],[")
            .map { it.replace('[', ' ').replace(']', ' ').trim().split(",") }
            .map { it.map { it[1] }.toCharArray() }.toTypedArray()
        solveSudoku(map)
//        println(map.map { it.joinToString(",") }.joinToString("\n"))
        println()
    }
}

//[3|5|6],[1|3|5],9,7,4,8,[1|6],[1|3|5],2
//7,[1|3|4|5|8],[1|3|5],6,[3|5],2,[1|4|8],[1|3|5|8],9
//[3|4|5|6|8],2,[3|5|6],1,[3|5],9,[4|6|7|8],[3|5|7|8],[3|4|5]
//[3|5],[3|5],7,9,8,6,2,4,1
//2,6,4,3,1,7,5,9,8
//1,9,8,5,2,4,3,6,7
//9,[1|4|5|7],[1|5],8,6,3,[1|4|7],2,[4|5]
//[3|5|8],[3|5|7|8],2,4,9,1,[7|8],[3|5|7|8],6
//[3|4|6|8],[1|3|4|8],[1|3|6],2,7,5,9,[1|3|8],[3|4]

//[3|5|6],[1|3|5],9,7,4,8,[1|6],[1|3|5],2
//7,[1|3|4|5|8],[1|3|5],6,[3|5],2,[1|4|8],[1|3|5|8],9
//[3|4|5|6|8],2,[3|5|6],1,[3|5],9,[4|6|7|8],[3|5|7|8],[3|4|5]
//[3|5],[3|5],7,9,8,6,2,4,1
//2,6,4,3,1,7,5,9,8
//1,9,8,5,2,4,3,6,7
//9,[1|4|5|7],[1|5],8,6,3,[1|4|7],2,[4|5]
//[3|5|8],[3|5|7|8],2,4,9,1,[7|8],[3|5|7|8],6
//[3|4|6|8],[1|3|4|8],[1|3|6],2,7,5,9,[1|3|8],[3|4]


