package main

type MyCircularDeque0 struct {
	size  int
	len   int
	head  int
	tail  int
	queue []int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor0(k int) MyCircularDeque0 {
	return MyCircularDeque0{
		size:  k,
		queue: make([]int, k),
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque0) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[this.head] = value
	this.len++
	this.head = (this.head + 1) % this.size
	if this.len == 1 {
		this.tail = (this.tail + this.size - 1) % this.size
	}
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque0) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[this.tail] = value
	this.len++
	this.tail = (this.tail + this.size - 1) % this.size
	if this.len == 1 {
		this.head = (this.head + 1) % this.size
	}
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque0) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + this.size - 1) % this.size
	this.len--
	if this.len == 0 {
		this.tail = (this.tail + 1) % this.size
	}
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque0) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail + 1) % this.size
	this.len--
	if this.len == 0 {
		this.head = (this.head + this.size - 1) % this.size
	}
	return true
}

func (this *MyCircularDeque0) GetFront() int {
	if this.len == 0 {
		return -1
	}
	return this.queue[(this.head+this.size-1)%this.size]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque0) GetRear() int {
	if this.len == 0 {
		return -1
	}
	return this.queue[(this.tail+1)%this.size]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque0) IsEmpty() bool {
	return this.len == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque0) IsFull() bool {
	return this.len == this.size
}
