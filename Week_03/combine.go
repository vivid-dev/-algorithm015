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
func combine(n int, k int) [][]int {
	res := [][]int{}
	backtrace(n, k, 1, []int{}, &res)
	return res
}

func backtrace(n, k, idx int, ans []int, res *[][]int) {
	if k == len(ans) {
		tmp := make([]int, k)
		copy(tmp, ans)
		*res = append(*res, tmp)

		return
	}

	for i := idx; i <= n; i++ {
		if len(ans) < k {
		    ans = append(ans, i)
			backtrace(n, k, i+ 1, ans, res)
			ans = ans[:len(ans)-1]
		} else {
			ans = ans[:len(ans)-1]
		}
	}
}