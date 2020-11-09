func findCircleNum(M [][]int) int {
    u := UnionSet{}
    u.Init(len(M))

    for i := 0; i < len(M); i++ {
        for j := i + 1; j < len(M[i]); j++ {
            if M[i][j] == 1 {
                u.MergeUnion(i, j)
            }
        }
    }

    return u.GetCnt()
}

// 并查集结构
type UnionSet struct {
    data []int
    cnt int
}

// 初始化并查集
func (u *UnionSet) Init(n int) {
    u.data = make([]int, n)
    u.cnt = n

    for i := 0; i < n; i++ {
        u.data[i] = i
    }
}

// 查找并查集根节点 & 进行树压缩
func (u *UnionSet) FindRoot(i int) int {
    for i != u.data[i] {
        u.data[i] = u.data[u.data[i]]
        i = u.data[i]
    }
    return i
}

// 合并两个并查集
func (u *UnionSet) MergeUnion(i, j int) {
    iRoot := u.FindRoot(i)
    jRoot := u.FindRoot(j)
    if iRoot == jRoot {
        return
    }
    u.data[iRoot] = jRoot
    u.cnt--
}

func (u UnionSet) GetCnt() int {
    return u.cnt
}