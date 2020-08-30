// 双指针法
// 思路：始终移动高度较小的一个边界，累加雨水总量
// i 从左向右迭代
// j 从右向左迭代
// left 左边界
// right 右边界
// min 记录边界较小值
// res 记录最终的总蓄水量
func trap(height []int) int {
    i, j := 0, len(height) - 1
    res, min, left, right := 0, 0, 0, 0
    
    for j > i {
        if height[i] > left {
            left = height[i]
        }

        if height[j] > right {
            right = height[j]
        }
        
        // 过滤掉无需处理的情况，因为边界为零没法存水
        if left == 0 {
            i++
            continue
        }
        if right == 0 {
            j--
            continue
        }

        min = getMin(left, right)

        if height[i] > height[j] {    // 需要挪动右边界
            res += min - height[j]
            j--
        } else {                      // 需要挪动左边界
            res += min - height[i]
            i++
        }
    }

    return res
}

func getMin(x, y int) int {
    if x > y {
        return y
    }
    return x
}