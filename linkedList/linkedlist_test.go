package linkedlist

import "testing"

func TestPush(t *testing.T) {
	h := new(Head)
	testData := []int{1, 2, 3, 4, 5, 6}

	for _, n := range testData {
		h.Push(n)
	}

	n := h.next
	for i := 0; n != nil; i++ {

		if n.value != testData[i] {
			t.Errorf("%d != %d", n.value, testData[i])
		}

		n = n.next
	}
}

func TestLen(t *testing.T) {
	h := new(Head)
	testData := []int{1, 2, 3, 4, 5, 6}

	for _, n := range testData {
		h.Push(n)
	}

	if h.Len() != len(testData) {
		t.Errorf("%d != %d", h.Len(), len(testData))
	}
}

func TestDelete(t *testing.T) {
	h := new(Head)
	testData := []int{1, 2, 3, 4, 5, 6}

	for _, n := range testData {
		h.Push(n)
	}

	n := h.next
	for n != nil {
		if !h.Delete(n) {
			t.Error("Delete failed")
		}
		n = n.next
	}

	if h.Len() != 0 {
		t.Errorf("Delete failed, len returned %d not 0", h.Len())
	}

	n = h.next
	for n != nil {
		if h.Delete(n) {
			t.Error("Delete failed, returned true when searching for deleted item")
		}
		n = n.next
	}

	if h.tail != nil {
		t.Error("Tail non nil after being deleted")
	}
}
