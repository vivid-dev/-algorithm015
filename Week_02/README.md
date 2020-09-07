### 提交说明
代码和题解思路都在代码文件里，方便阅读。<br>
### 本周收获
1.了解了树形数据结构<br>
2.了解并练习了树形相关算法和解题思路<br>
3.了解了golang堆的实现和使用<br>
4.了解并练习了dfs算法<br>
### golang heap(优先队列)实现和使用
1.要实现的接口<br>
```
type Elem struct {
	Val      int
	Priority int
}

type IntHeap []Elem

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority
}

func (h IntHeap) Peek() Elem {
    return h[0]
}

func (h *IntHeap) Push(val interface{}) {
	*h = append(*h, val.(Elem))
}

func (h *IntHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}
```
2.使用方法<br>
```
初始化:
h := &IntHeap{}
heap.Init(h)
增加元素:
heap.Push(h, Elem{Val: key, Priority: v})
弹出元素:
heap.Pop(h)
计算长度:
h.Len()
```
