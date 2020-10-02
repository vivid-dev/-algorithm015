// 解题思路:二分法
// 时间复杂度:O(logN) 空间复杂度:O(1)
func mySqrt(x int) int {
    return search(0, x, x)
}

func search(l, r, x int) int {
    if l > r {
        return r
    }

    mid := l + (r - l) >> 1
    s := mid * mid
    if s == x {
        return mid
    } else if s < x {
        l = mid + 1
    } else {
        r = mid - 1
    }
    return search(l, r, x)
}