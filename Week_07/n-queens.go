// 解题思路：并查集
// 时间复杂度O(2^N), 空间复杂度O(N^2)
func solveNQueens(n int) [][]string {
	res := [][]string{}
	backtrace(n, []string{}, map[int]bool{}, map[int]bool{}, map[int]bool{}, &res)

	return res
}

func backtrace(cnt int, ans []string, s, p, n map[int]bool, res *[][]string) {
	// 出口条件
	if cnt == len(ans) {
		tmp := make([]string, len(ans))
		copy(tmp, ans)
		*res = append(*res, tmp)
		return
	}

	// 筛选
	for i := 0; i < cnt; i++ {
		// pie na map 为空
		sOk, _ := s[i]
		pOk, _ := p[len(ans)+i]
		nOk, _ := n[len(ans)-i]
		if sOk || pOk || nOk {
			continue
		}
		// 选中
		str := make([]byte, cnt)
		for j := 0; j < cnt; j++ {
			str[j] = byte('.')
		}

		str[i] = byte('Q')
		s[i] = true
		p[len(ans)+i] = true
		n[len(ans)-i] = true
		ans = append(ans, string(str))

		//backtrace
		backtrace(cnt, ans, s, p, n, res)
		//revert
		ans = ans[:len(ans)-1]
		str[i] = byte('.')
		delete(s, i)
		delete(p, len(ans)+i)
		delete(n, len(ans)-i)
	}
}