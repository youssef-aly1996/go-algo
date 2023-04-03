package linkedlist

type linkedListNode struct {
	data int
	next *linkedListNode
}

type LinkedList struct {
	Head *linkedListNode
	Tail *linkedListNode
}

type LinkedListIterator struct {
	current *linkedListNode
}

func (li LinkedListIterator) data() int {
	return li.current.data
}

func (li LinkedListIterator) next() {
	li.current = li.current.next
}

func linkedListConstructor(_data int) *linkedListNode {
	return &linkedListNode{data: _data, next: nil}
}

func (ls *LinkedList) New(_data int) {
	newNode := linkedListConstructor(_data)
	if ls.Head == nil {
		ls.Head = newNode
		ls.Tail = newNode
	} else {
		ls.Tail.next = newNode
		ls.Tail = newNode
	}
}
