// 双指针解法
// 思路: 删掉要删的实际是留下要留的
// j 标识最后一个非重复元素的下标
// i 从头开始向后比较, 填j+1对应的位置
// 所以回包应该返回j+1
func removeDuplicates(nums []int) int {
    j := 0
    for i := range nums {
        if nums[i] != nums[j] && i > j{
            nums[i], nums[j+1] = nums[j+1], nums[i]
            j++
        }
    }

    return j + 1 
}
