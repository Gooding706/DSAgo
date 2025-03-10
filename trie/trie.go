package trie

type Trie struct {
	value       any
	connections map[rune]*Trie
}

func (t *Trie) Insert(key string, val any) bool {
	node := t
	for _, char := range key {
		next := node.connections[char]
		if next == nil {
			if node.connections == nil {
				node.connections = map[rune]*Trie{}
			}
			next = new(Trie)
			node.connections[char] = next
		}
		node = next
	}

	isNew := node.value == nil
	node.value = val

	return isNew
}

func (t *Trie) Get(key string) any {
	node := t
	for _, char := range key {
		next := node.connections[char]
		if next == nil {
			return nil
		}

		node = next
	}

	return node.value
}

func (t *Trie) Delete(key string) bool {
	path := make([]*Trie, len(key)+1)
	path[0] = t

	node := t
	for i, char := range key {
		next := node.connections[char]
		if next == nil {
			return false
		}

		node = next
		path[i+1] = node
	}

	path[len(key)].value = nil

	for i := len(key); i > 0; i-- {
		char := []rune(key)[i-1]
		if len(path[i].connections) > 0 || path[i].value != nil {
			break
		}
		delete(path[i-1].connections, char)
	}
	return true
}
