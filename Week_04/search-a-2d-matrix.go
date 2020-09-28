/*
// 解题思路1：先遍历找到对应行,再对列进行二分查找,时间复杂度O(M + log(N)),空间复杂度:O(1)
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    if m == 0 {
        return false
    }
    n := len(matrix[0])
    if n == 0 {
        return false
    }

    idx := -1
    for i := range matrix {
        if matrix[i][0] <= target && target <= matrix[i][n-1] {
            idx = i
        }
    }

    if idx == -1 {
        return false
    }

    return BinarySearch(matrix[idx], target)
}

func BinarySearch(nums []int, target int) bool {
    if len(nums) == 0 {
        return false
    }

    mid := (len(nums)-1) >> 1
    if nums[mid] == target {
        return true
    } else if nums[mid] > target {
        return BinarySearch(nums[:mid], target)
    } else {
        return BinarySearch(nums[mid+1:], target)
    }
}
*/
// 解题思路2：二维转一维，进行二分查找,时间复杂度O(log(N+M)),空间复杂度:O(1)
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    if m == 0 {
        return false
    }
    n := len(matrix[0])
    if n == 0 {
        return false
    }
    return BinarySearch(matrix, 0, m*n-1, n, target)
}

func BinarySearch(matrix [][]int, l, r, n, t int) bool {
    if l > r {
        return false
    }

    mid := l + (r - l) >> 1
   
    i := mid / n
    j := mid % n

    if matrix[i][j] == t {
        return true
    } else if matrix[i][j] > t {
        return BinarySearch(matrix, l, mid-1, n, t)
    } else {
        return BinarySearch(matrix, mid+1, r, n, t)
    }
}