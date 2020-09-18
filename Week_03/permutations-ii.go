func permuteUnique(nums []int) [][]int {
    res := make([][]int, 0)
    backtrack(nums, []int{}, map[int]bool{}, &res)
    return res
}

func backtrack(nums, ans []int, filter map[int]bool, res *[][]int) {
    // 退出条件
    if len(nums) == len(ans) {
        tmp := make([]int, len(ans))
        copy(tmp, ans)
        *res = append(*res, tmp)
    }
    repeate := make(map[int]bool, 0)
    // 进行backtrack处理
    for i := range nums {
        // 选择处理
        if _, ok := filter[i]; ok {
            continue
        }
        if _, ok := repeate[nums[i]]; ok {
            continue
        }
        // 添加结果集合
        filter[i] = true
        repeate[nums[i]] = true
        ans = append(ans, nums[i])

        // backtrack
        backtrack(nums, ans, filter, res)
        // revert
        ans = ans[:len(ans)-1]
        delete(filter, i)
    }
}