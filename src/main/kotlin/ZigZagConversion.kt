fun convert(s: String, numRows: Int): String {
    if (numRows == 1) return s
    val list = Array<MutableList<Char>>(numRows) { ArrayList() }
    s.forEachIndexed { index, char ->
        val i = index % (numRows - 1)
        val j = index / (numRows - 1)
        val p = if (j % 2 == 0) {
            i
        } else {
            (numRows - 1) - i
        }
        list[p].add(char)
    }
    val sb = StringBuffer()
    for (chars in list) {
        for (char in chars) {
            sb.append(char)
        }
    }
    return sb.toString()
}

fun convert2(s: String, numRows: Int): String {
    if (s.length <= numRows) return s
    val sb = StringBuffer(s)
    val len = IntArray(numRows)

    val iterateLen = (numRows - 1) * 2 //迭代的长度
    val iterateSize = s.length / iterateLen //迭代次数
    //第一个和最后一个是size，其他是2*size
    len[0] = iterateSize
    for (x in 1 until (numRows - 1)) {
        len[x] = iterateSize * 2
    }
    len[numRows - 1] = iterateSize
    //对多出的部分所占的空间进行处理
    val mod = s.length % iterateLen //余出的部分
    for (y in 1 until mod + 1) { //只有余数大于1的才有操作
        if (y < numRows) {
            len[y - 1]++
        } else {
            len[2 * numRows - 1 - y]++
        }
    }
    //计算每个列表开始的下标
    var temp = s.length
    for (x in (0 until numRows).reversed()) {
        temp -= len[x]
        len[x] = temp
    }
    //开始写入
    s.forEachIndexed { index, c ->
        val i = index % (numRows - 1)
        val j = index / (numRows - 1)
        val loc = if (j % 2 == 0) {
            i
        } else {
            (numRows - 1) - i
        }
        val newIndex = len[loc]++
        sb.setCharAt(newIndex, c)
    }
    return sb.toString()
}

fun convert3(s: String, numRows: Int): String {
    if (s.length <= numRows) return s
    val length = s.length
    val ret = StringBuffer()
    val cycleLen = 2 * numRows - 2
    for (i in 0 until numRows) {
        var j = 0
        while (j + i < length) {
            ret.append(s[i + j])
            //不是第一行，不是最后一行
            if (i != 0 && i != numRows - 1 && j + cycleLen - i < length) {
                ret.append(j + cycleLen - i)
            }
            j += cycleLen
        }
    }
    return ret.toString()
}

fun main() {
//    println(convert("PAYPALISHIRING", 3))
//    println(convert2("PAYPALISHIRING", 3))
    println(convert2("ABC", 3))
}