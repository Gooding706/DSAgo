package trie

import "testing"

func testReinsert(t *testing.T, tr *Trie) {
	expectations := map[string]int{
		"foo":      1,
		"bar":      2,
		"food":     3,
		"bard":     4,
		"barstool": 5}

	for k, v := range expectations {
		b := tr.Insert(k, v)
		if b {
			t.Errorf("Expected insert of existint key to return false, trie.Insert(%s, ?) = true", k)
		}
	}
}

func testGet(t *testing.T, tr *Trie) {
	expectations := map[string]int{
		"foo":      1,
		"bar":      2,
		"food":     3,
		"bard":     4,
		"barstool": 5}

	for k, v := range expectations {
		ret := tr.Get(k)
		if ret != v {
			t.Errorf("Incorrect value returned for key: %s", k)
		}
	}
}

func testDelete(t *testing.T, tr *Trie) {
	expectations := map[string]int{
		"foo":      1,
		"bar":      2,
		"food":     3,
		"bard":     4,
		"barstool": 5}

	for k := range expectations {
		b := tr.Delete(k)
		if !b {
			t.Errorf("Delete should return true if key %s exists", k)
		}
	}

	for k := range expectations {
		ret := tr.Get(k)
		if ret != nil {
			t.Errorf("Deleted value should return nil, key: %s", k)
		}
	}
}

func TestTrie(t *testing.T) {
	tr := new(Trie)
	expectations := map[string]int{
		"foo":      1,
		"bar":      2,
		"food":     3,
		"bard":     4,
		"barstool": 5}

	for k, v := range expectations {
		b := tr.Insert(k, v)
		if !b {
			t.Errorf("Expected insert of non-existant key to return true, trie.Insert(%s, ?) = false", k)
		}
	}

	testReinsert(t, tr)

	testGet(t, tr)

	testDelete(t, tr)
}
