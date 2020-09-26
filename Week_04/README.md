### 提交说明
代码和题解思路都在代码文件里，方便阅读。<br>
### 本周收获
1.学习并练习了搜索算法，dfs和bfs。<br>
2.学习并练习了贪心算法。<br>
3.学习并练习了分治算法，二分查找的递归和非递归实现如下，复杂度：O(logN)<br>
```
// 递归实现
func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    return helper(nums, l, r, target)
}

func helper(nums []int, l, r, t int) int {
    if l > r {
        return -1
    }

    mid := l + (r - l) >> 1
    if nums[mid] == t {
        return mid
    } 
    if nums[mid] > t {
        r = mid - 1
    } else {
        l = mid + 1
    }
   
    return helper(nums, l, r, t)
}

// 循环实现
func search(nums []int, target int) int {
    l, r, mid := 0, len(nums) - 1, 0
    
    for l <= r {
        mid = l + (r - l) >> 1
        if nums[mid] == target {
            return mid
        }
        if nums[mid] > target {
            r = mid - 1
        } else  {
            l = mid + 1
        }
    }

    return -1
}
```