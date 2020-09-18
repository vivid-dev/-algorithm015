// 回溯法
// 思路:回溯模板
/*
回溯模板：
result = []
func backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return
    
    for 选择 range 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
*/
func permute(nums []int) [][]int {
    list := [][]int{}

    backtrace(nums, []int{}, map[int]bool{}, &list)

    return list
}

func backtrace(nums []int, ans []int, filter map[int]bool, list *[][]int) {
    // 如果满足条件添加答案到结果集合
    if len(nums) == len(ans) {
        tmp := make([]int, len(ans))
        copy(tmp, ans)
        *list = append(*list, tmp)
    }
    
    // 遍历选择路径做处理
    for i := 0; i < len(nums); i++  {
        // 做选择
        if _, ok := filter[i]; ok {
            continue
        }
        filter[i] = true

        ans = append(ans, nums[i])
        backtrace(nums, ans, filter, list)
        ans = ans[:len(ans) - 1]
        delete(filter, i)
    }
}