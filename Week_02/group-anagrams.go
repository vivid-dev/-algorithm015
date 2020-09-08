// 查表法
// 解题思路：分别给各个元素排序, 并存储在map里，方便查找归类
// tmp 用于存储排序后的各个元素
// res 用于存储最终结果
// m 存储每个元素在res里的位置
// 复杂度:时间复杂度:O(N*MlogM) 空间复杂度:O(N+M)
func groupAnagrams(strs []string) [][]string {
    tmp := make([]string, len(strs))
    res := make([][]string, 0)
    copy(tmp, strs)
    m := make(map[string]int, 0)
    for i := range strs {
        s := []byte(tmp[i])
        sort.Slice(s, func(a, b int) bool {return s[a] < s[b]})
        tmp[i] = string(s)
        if v, ok := m[tmp[i]]; ok {
            res[v] = append(res[v], strs[i])
        } else {
            res = append(res, []string{strs[i]})
            m[tmp[i]] = len(res) - 1
        }
    }

    return res
}