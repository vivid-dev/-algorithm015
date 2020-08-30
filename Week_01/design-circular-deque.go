type MyCircularDeque struct {
    Size int
    Curr int
    Front *MyCircularDequeNode
    Tail *MyCircularDequeNode
}

type MyCircularDequeNode struct {
    Prev *MyCircularDequeNode
    Next *MyCircularDequeNode
    Val int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    head := MyCircularDeque{}
    head.Size = k
    head.Curr = 0

    return head
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }

    node := MyCircularDequeNode{}
    node.Val = value
    if !this.IsEmpty() {
        node.Prev = this.Front.Prev
        node.Next = this.Front
        this.Front.Prev = &node
		this.Tail.Next = &node
		this.Front = &node
    } else {
        this.Front = &node
        this.Tail = &node
        node.Prev = &node
        node.Next = &node
    }
    
    this.Curr++
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    node := MyCircularDequeNode{}
    node.Val = value
    if !this.IsEmpty() {
        node.Next = this.Tail.Next
        node.Prev = this.Tail
        this.Tail.Next = &node
		this.Front.Prev = &node
		this.Tail = &node
    } else {
        this.Front = &node
        this.Tail = &node
        node.Prev = &node
        node.Next = &node
    }
    
    this.Curr++
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    this.Curr--
    if this.IsEmpty() {
        this.Tail = nil
        this.Front = nil
    } else {
        this.Front = this.Front.Next
        this.Front.Prev = this.Tail
        this.Tail.Next = this.Front
    }
    
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    this.Curr--
    if this.IsEmpty() {
        this.Tail = nil
        this.Front = nil
    } else {
        this.Tail = this.Tail.Prev
        this.Tail.Next = this.Front
        this.Front.Prev = this.Tail
    }
    
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.Front.Val
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.Tail.Val
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    if this.Curr == 0 {
        return true
    }
    return false
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    if this.Size == this.Curr {
        return true
    } 
    return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */