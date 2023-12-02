package main

type trieNode struct {
	value byte
	leaf  bool
	chars map[byte]*trieNode
}

func NewTrieNode(value ...byte) *trieNode {
	var v byte
	if len(value) != 0 {
		v = value[0]
	} else {
		v = 0
	}
	return &trieNode{value: v, chars: make(map[byte]*trieNode)}
}

func (t *trieNode) insert(s string) {
	node := t

	for _, sc := range s {
		c := byte(sc)
		if _, ok := node.chars[c]; !ok {
			node.chars[c] = NewTrieNode(c)
		}
		node = node.chars[c]
	}
	node.leaf = true
}

func (t *trieNode) find(s string) bool {
	node := t
	for _, char := range s {
		node = node.chars[byte(char)]
		if node == nil {
			return false
		}
	}
	return node.leaf
}

func (t *trieNode) findSubstr(s string) bool {
	// Tries (heh) to find a substring of in the trie
	// i.e., if s is "foobar", and the trie contains "foo",
	// the function will return true
	node := t
	for _, char := range s {
		node = node.chars[byte(char)]
		if node == nil {
			return false
		}
		if node.leaf {
			return true
		}
	}
	return node.leaf

}
