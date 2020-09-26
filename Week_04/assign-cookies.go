func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    i := 0
    j := 0
    cnt := 0

    for i < len(g) && j < len(s) {
        if g[i] <= s[j] {
            cnt++
            i++
        }
        j++
        
    }
    
    return cnt
}