// 双指针法
// 思路：要把0放到后面相当于把非零元素放到前面
// i 迭代nums, 找非0元素，来和
// j 要交换的位置。
func moveZeroes(nums []int)  {
    j := 0
    for i := range nums {
        if nums[i] != 0 {
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
    }
}
