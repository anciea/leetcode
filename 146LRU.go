package main

type LRUItem struct {
	key   int
	value int

	pre  *LRUItem
	next *LRUItem
}

type LRUCache struct {
	vis      map[int]*LRUItem
	head     *LRUItem  // 最旧元素
	tail     *LRUItem  // 最新元素
	size     int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		vis:      make(map[int]*LRUItem),
		head:     nil,
		tail:     nil,
		size:     0,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if item, ok := this.vis[key]; ok {
		ans := item.value
		this.moveToTail(item)
		return ans
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if item, ok := this.vis[key]; ok {
		// Existing key: update value and move to tail
		item.value = value
		this.moveToTail(item)
	} else {
		// New key: create new node and add to tail
		newNode := &LRUItem{
			key:   key,
			value: value,
			pre:   nil,
			next:  nil,
		}
		this.vis[key] = newNode
		this.addToTail(newNode)
	}

	// Check capacity and remove head if necessary
	if this.size > this.capacity {
		this.removeHead()
	}
}

func (this *LRUCache) addToTail(node *LRUItem) {
	if this.head == nil {
		// Empty list
		this.head = node
		this.tail = node
		node.pre = nil
		node.next = nil
	} else {
		// Add to tail
		this.tail.next = node
		node.pre = this.tail
		node.next = nil
		this.tail = node
	}
	this.size++
}

func (this *LRUCache) removeNode(node *LRUItem) {
	if node.pre != nil {
		node.pre.next = node.next
	} else {
		// node is head
		this.head = node.next
	}

	if node.next != nil {
		node.next.pre = node.pre
	} else {
		// node is tail
		this.tail = node.pre
	}

	// Clean up node pointers
	node.pre = nil
	node.next = nil
	this.size--
}

func (this *LRUCache) moveToTail(node *LRUItem) {
	this.removeNode(node)
	this.addToTail(node)
}

func (this *LRUCache) removeHead() {
	if this.head != nil {
		oldHead := this.head
		delete(this.vis, oldHead.key)
		this.removeNode(oldHead)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
