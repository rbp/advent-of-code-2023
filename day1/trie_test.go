package main

import "testing"

func TestTrieEmpty(t *testing.T) {
	trie := NewTrieNode()

	if len(trie.chars) != 0 {
		t.Fatal("Trie chars is not empty")
	}
}

func TestTrieOneElement(t *testing.T) {
	trie := NewTrieNode()

	trie.insert("foo")
	if !(len(trie.chars) == 1 &&
		trie.chars['f'] != nil) {
		t.Fatal("Trie root is invalid")
	}
	f := trie.chars['f']
	if !(len(f.chars) == 1 &&
		f.chars['o'] != nil &&
		f.leaf == false) {
		t.Fatal("Trie['f'] is invalid")
	}
	o1 := f.chars['o']
	if !(len(o1.chars) == 1 &&
		o1.chars['o'] != nil &&
		o1.leaf == false) {
		t.Fatal("Trie['f']['o'] is invalid")
	}
	o2 := o1.chars['o']
	if !(len(o2.chars) == 0 &&
		o2.leaf == true) {
		t.Fatalf("Trie['f']['o']['o'] is invalid (%v %v)", o2.chars, o2.leaf)
	}
}

func TestTrieOverlappingElements(t *testing.T) {
	trie := NewTrieNode()

	trie.insert("foo")
	trie.insert("fold")
	trie.insert("fo")

	if !(len(trie.chars) == 1 &&
		trie.chars['f'] != nil) {
		t.Fatal("Trie root is invalid")
	}
	f := trie.chars['f']
	if !(len(f.chars) == 1 &&
		f.chars['o'] != nil &&
		f.leaf == false) {
		t.Fatal("Trie['f'] is invalid")
	}
	o1 := f.chars['o']
	if !(len(o1.chars) == 2 &&
		o1.chars['o'] != nil &&
		o1.chars['l'] != nil &&
		o1.leaf == true) {
		t.Fatal("Trie['f']['o'] is invalid")
	}
	o2 := o1.chars['o']
	if !(len(o2.chars) == 0 &&
		o2.leaf == true) {
		t.Fatal("Trie['f']['o']['o'] is invalid")
	}

	l := o1.chars['l']
	if !(len(l.chars) == 1 &&
		l.chars['d'] != nil &&
		l.leaf == false) {
		t.Fatal("Trie['f']['o']['l'] is invalid")
	}
	d := l.chars['d']
	if !(len(d.chars) == 0 &&
		d.leaf == true) {
		t.Fatal("Trie['f']['o']['o'] is invalid")
	}

}

func TestTrieNonOverlapping(t *testing.T) {
	trie := NewTrieNode()

	trie.insert("foo")
	trie.insert("fold")
	trie.insert("new")

	if !(len(trie.chars) == 2 &&
		trie.chars['f'] != nil &&
		trie.chars['n'] != nil) {
		t.Fatal("Trie root is invalid")
	}
	f := trie.chars['f']
	if !(len(f.chars) == 1 &&
		f.chars['o'] != nil &&
		f.leaf == false) {
		t.Fatal("Trie['f'] is invalid")
	}
	o1 := f.chars['o']
	if !(len(o1.chars) == 2 &&
		o1.chars['o'] != nil &&
		o1.chars['l'] != nil &&
		o1.leaf == false) {
		t.Fatal("Trie['f']['o'] is invalid")
	}
	o2 := o1.chars['o']
	if !(len(o2.chars) == 0 &&
		o2.leaf == true) {
		t.Fatal("Trie['f']['o']['o'] is invalid")
	}

	n := trie.chars['n']
	if !(len(n.chars) == 1 &&
		n.chars['e'] != nil &&
		n.leaf == false) {
		t.Fatal("Trie['n'] is invalid")
	}
	e := n.chars['e']
	if !(len(e.chars) == 1 &&
		e.chars['w'] != nil &&
		e.leaf == false) {
		t.Fatal("Trie['n']['e'] is invalid")
	}
	w := e.chars['w']
	if !(len(w.chars) == 0 &&
		w.leaf == true) {
		t.Fatal("Trie['n']['e']['w'] is invalid")
	}

}

func TestTrieFind(t *testing.T) {
	trie := NewTrieNode()

	trie.insert("foo")
	trie.insert("fold")
	trie.insert("new")

	if !trie.find("foo") {
		t.Fatal("Trie cannot find 'foo'")
	}
	if !trie.find("fold") {
		t.Fatal("Trie cannot find 'fold'")
	}
	if !trie.find("new") {
		t.Fatal("Trie cannot find 'new'")
	}

	if trie.find("bar") {
		t.Fatal("Trie found 'bar'")
	}
	if trie.find("") {
		t.Fatal("Trie found ''")
	}
	if trie.find("fo") {
		t.Fatal("Trie found 'fo'")
	}
}

func TestTrieFindSubstr(t *testing.T) {
	trie := NewTrieNode()

	trie.insert("foo")
	trie.insert("fold")
	trie.insert("new")

	if !trie.findSubstr("foobar") {
		t.Fatal("Trie cannot find 'foo' in 'foobar'")
	}
	if !trie.find("new") {
		t.Fatal("Trie cannot find 'new' in 'newly'")
	}
}
