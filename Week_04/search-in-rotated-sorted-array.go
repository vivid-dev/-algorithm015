func search(nums []int, target int) int {
    return helper(nums, 0, len(nums) - 1, target)
}

func helper(nums []int, l, r, t int) int {
    if l > r {
        return -1
    }

    mid := l + (r - l) >> 1
    //fmt.Println(mid)
    if nums[mid] == t {
        return mid
    }

    // 通过边界值判断中间区间是不是顺序单调，失序只可能出现在一边
    if nums[l] <= nums[mid]   {  // 左半边单调增
        if nums[l] <= t && t < nums[mid] {  // 在左边找
            r = mid - 1
        } else {    // 在右边找
            l = mid + 1
        }
    } else {    // 右边单调增
        if nums[mid] < t && t <= nums[r] {   // 在右边找
            l = mid + 1
        } else {  // 在左边找
            r = mid - 1
        }
    }

   return helper(nums, l, r, t)
}