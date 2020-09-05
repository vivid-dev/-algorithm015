// 查表法
// 思路:把s存到map ms中,遍历t判断是否都在t中
// i 用来遍历s和t
// ms里存s
// flag存储最终结果
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    flag := true
    ms := make(map[byte]int, 0)
    for i := range(s) {
        ms[s[i]]++
    }

    for i := range(t) {
        if v, ok := ms[t[i]]; ok && v != 0 {
            ms[t[i]]--
        } else {
            flag = false
            break
        }
    }

    return flag
}