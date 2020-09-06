// 小顶堆(优先队列)解法 超过100%
// 思路: 先将数据存入map，再根据优先队列取top k
// m 用于存储和计数
// h 小顶堆(优先队列)


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

func (h *IntHeap) Push(val interface{}) {
	*h = append(*h, val.(Elem))
}

func (h *IntHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for i := range nums {
		m[nums[i]]++
	}
	h := &IntHeap{}
	heap.Init(h)

	for key, v := range m {
		if h.Len() < k {
			heap.Push(h, Elem{Val: key, Priority: v})
		} else {
			tmp := heap.Pop(h)
			if tmp.(Elem).Priority < v {
				heap.Push(h, Elem{Val: key, Priority: v})
			} else {
				heap.Push(h, tmp)
			}
		}
	}
    
	ans := []int{}
	for i := 0; i < k; i++ {
		tmp := heap.Pop(h)
		ans = append(ans, tmp.(Elem).Val)
	}

	return ans
}
