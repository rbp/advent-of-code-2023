package main

type Mapunit struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

func (m *Mapunit) srcMatches(src int) bool {
	return src >= m.sourceStart && src <= m.sourceStart+m.rangeLen
}

type Catmap struct {
	maps    []*Mapunit
	maptree *MapTreeNode
}

func NewCatmap(maps []*Mapunit) *Catmap {
	catmap := Catmap{maps: maps}

	catmap.maptree = &MapTreeNode{mapunit: maps[0]}
	for _, mapunit := range maps[1:] {
		catmap.maptree.add(
			&MapTreeNode{mapunit: mapunit},
		)
	}
	return &catmap
}

func (c *Catmap) find(src int) int {
	dst := c.maptree.find(src)
	if dst >= 0 {
		return dst
	}
	// Business rule: if there's no mapping, dst = src
	return src
}

type MapTreeNode struct {
	mapunit *Mapunit
	left    *MapTreeNode
	right   *MapTreeNode
}

func (tree *MapTreeNode) add(node *MapTreeNode) {
	if node.mapunit.sourceStart < tree.mapunit.sourceStart {
		if tree.left == nil {
			tree.left = node
		} else {
			tree.left.add(node)
		}
	} else if node.mapunit.sourceStart > tree.mapunit.sourceStart {
		if tree.right == nil {
			tree.right = node
		} else {
			tree.right.add(node)
		}
	} else {
		panic("Trying to insert an equal map??")
	}
}

// FindMapping returns the correct destination for the given source
func (tree *MapTreeNode) find(src int) int {
	if tree.mapunit.srcMatches(src) {
		r := src - tree.mapunit.sourceStart
		return tree.mapunit.destStart + r
	}
	if tree.left != nil && src < tree.mapunit.sourceStart {
		return tree.left.find(src)
	}
	if tree.right != nil && src > tree.mapunit.sourceStart {
		return tree.right.find(src)
	}
	return -1
}
