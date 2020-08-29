// 循环迭代解法
// 思路: 存入map中，方便查找
func twoSum(nums []int, target int) []int {
    m := make(map[int]int, 0)
    for k, v := range nums {
        if tmp, ok := m[target - v]; ok {
            return []int{tmp, k}
        }
        m[v] = k
    }

    return []int{}
}
