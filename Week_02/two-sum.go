// 循环迭代解法
// 思路: 转换a+b=c,为枚举a, 寻找是否存在-c, 引入map，方便查找
// k, v迭代可选集合

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