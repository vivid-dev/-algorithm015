// 循环迭代法
// 思路：模拟列竖式计算加法
// i 从后往前用来迭代处理digits的元素来处理加法和进位
// cp 进位
func plusOne(digits []int) []int {
    // 因为是加一可以考虑为最低位(个位）的进位为1
    cp := 1
    for i := len(digits) - 1; i >= 0; i-- {
        digits[i] = digits[i] + cp
        cp = digits[i] / 10
        digits[i] = digits[i] % 10 
    }

    // 如果cp不为零，需要做下cp和当前结果的连接
    if cp != 0 {
        digits = append([]int{cp}, digits...)
    }

    return digits
}