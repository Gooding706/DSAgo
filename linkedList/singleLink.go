package linkedlist

type Head struct {
	next *SingleNode
	tail *SingleNode
}

type SingleNode struct {
	next  *SingleNode
	value any
}

func NewSinglyLinkedList() *Head {
	out := new(Head)
	out.next = nil
	out.tail = out.next
	return out
}

func newListNode(v any) *SingleNode {
	newNode := new(SingleNode)
	newNode.value = v
	newNode.next = nil
	return newNode
}

func (h *Head) Push(v any) {
	nextSpace := &h.next
	if h.tail != nil {
		nextSpace = &(h.tail.next)
	}

	newNode := newListNode(v)
	*nextSpace = newNode
	h.tail = newNode
}

func (h *Head) Delete(n *SingleNode) bool {
	neighbor := h.next

	if neighbor == n {
		h.next = n.next
		if h.tail == n {
			h.tail = nil
		}
		return true
	}

	for neighbor != n && neighbor != nil {
		neighbor = neighbor.next
	}

	if neighbor == nil {
		return false
	}

	neighbor.next = n.next
	if h.tail == n {
		h.tail = neighbor
	}

	return true
}

func (h *Head) Len() int {
	var len int = 0
	n := h.next
	for n != nil {
		len++
		n = n.next
	}
	return len
}
