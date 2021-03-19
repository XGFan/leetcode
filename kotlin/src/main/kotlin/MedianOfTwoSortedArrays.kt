fun findMedianSortedArrays(nums1: IntArray, nums2: IntArray): Double {
    var m = nums1
    var n = nums2
    if (m.size > n.size) {
        val temp = m
        m = n
        n = temp
    }
    val ml = m.size
    val nl = n.size
    //到此，m长度小于n的长度
    val halfLen = (ml + nl + 1) / 2 //中间点或者，中间偏左边的点
    var iMin = 0
    var iMax = ml
    while (iMin <= iMax) {
        val i = (iMin + iMax) / 2
        val j = halfLen - i //这是为了始终保持左右两边长度一致
        if (i < iMax && m[i] <= n[j - 1]) { //m[i]太小，需要往右边二分查找
            iMin = i + 1
        } else if (i > iMin && m[i - 1] > n[j]) { //m[i]太大，需要往左边二分查找
            iMax = i - 1
        } else {
            //现在要么i到头了(最左端/最右端)，要么 m[i] > n[j-1] && m[i-1] <= n[j]
            //并且,i,j左右两边长度之和相等。
            //现在i左边有i个数，右边有ms-i-1个数
            //j = (ml + nl + 1) / 2 - i
            //j 左边有
            val maxLeft = when {
                i == 0 -> n[j - 1] //如果i在最左端，
                j == 0 -> m[i - 1]//这是m,n一样长,并且i在最右端
                else -> maxOf(m[i - 1], n[j - 1])
            }

            if ((ml + nl) % 2 == 1) { //只需要取一个值
                return maxLeft.toDouble()
            }

            val minRight = when {
                i == ml -> n[j]
                j == nl -> m[i]
                else -> minOf(n[j], m[i])
            }
            return (maxLeft + minRight) / 2.0;
        }
    }
    return 0.0;
}
//长度y的数组，下标x左边有x个数，右边有y-x-1个数

fun main() {

}