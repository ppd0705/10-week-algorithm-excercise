package main

type DNode struct {
	key   int
	value int
	prev  *DNode
	next  *DNode
}

type LRUCache struct {
	cap   int
	root  *DNode
	cache map[int]*DNode
}

func Constructor(capacity int) LRUCache {
	root := &DNode{}
	root.prev = root
	root.next = root
	return LRUCache{
		cap:   capacity,
		root:  root,
		cache: map[int]*DNode{},
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.movetoEnd(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.movetoEnd(node)
	} else {
		if len(this.cache) < this.cap {
			node := &DNode{
				key:   key,
				value: value,
				prev:  this.root.prev,
				next:  this.root,
			}
			this.root.prev.next = node
			this.root.prev = node
			this.cache[key] = node
		} else {
			this.root.key, this.root.value = key, value
			this.cache[key] = this.root
			this.root = this.root.next
			delete(this.cache, this.root.key)
		}
	}
}

func (this *LRUCache) movetoEnd(node *DNode) {
	node.prev.next, node.next.prev = node.next, node.prev
	node.prev, node.next = this.root.prev, this.root
	this.root.prev.next, this.root.prev = node, node
}
