// 解题思路：贪心算法，排序后依次比较
// 时间复杂度:O(N), 空间复杂度O(1)
func lemonadeChange(bills []int) bool {
    cp5 := 0
    cp10 := 0

    flag := true
    
    for i := range bills {
        if bills[i] == 5 {
            cp5++
        }
        if bills[i] == 10 {
            if cp5 == 0 {
                flag = false
                break
            }
            cp5--
            cp10++
        }
        if bills[i] == 20 {
            if cp5 == 0 || (cp10 == 0 && cp5 < 3) {
                flag = false
                break
            }
            if cp10 == 0 {
                cp5 -= 3
            } else {
                cp5--
                cp10--
            }
        }
    }

    return flag
}