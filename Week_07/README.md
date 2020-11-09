### 提交说明
代码和题解思路都在代码文件里，方便阅读。<br>
### 本周收获
1.学习并练习了Trie。<br>
2.学习并练习了并查集。<br>
3.学习了高级搜索。<br>
4.学习了AVL树和红黑树。<br>

### 相关模板
1.trie树模板
```
type Trie struct {
    Data map[byte]Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    t :=  Trie{Data:make(map[byte]Trie)}
    t.Data['#'] = Trie{Data:make(map[byte]Trie)}
    return t
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    node := this.Data
    for i := range word {
        if v, ok := node[word[i]]; !ok {
            node[word[i]] = Trie{Data:make(map[byte]Trie)}
            node = node[word[i]].Data
        } else {
            node = v.Data
        }
    }
    node['#'] = Trie{Data:make(map[byte]Trie)}
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    node := this.Data
    for i := range word {
        if v, ok := node[word[i]]; !ok {
            fmt.Println("vivid")
            return false
        } else {
            node = v.Data
        }
    }
    if _, ok := node['#']; ok {
        return true
    }

    return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    node := this.Data
    for i := range prefix {
        if v, ok := node[prefix[i]]; !ok {
            return false
        } else {
            node = v.Data
        }
    }
    return true
}
```
2.并查集模板
```

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
```