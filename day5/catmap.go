package main

type mapping struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

func (m *mapping) containsSource(src int) bool {
	return src >= m.sourceStart && src < m.sourceStart+m.rangeLen
}

type categoryMap struct {
	maps    []*mapping
	maptree *mapTreeNode
}

func newCategoryMap(maps []*mapping) *categoryMap {
	catmap := categoryMap{maps: maps}

	catmap.maptree = &mapTreeNode{mapunit: maps[0]}
	for _, mapunit := range maps[1:] {
		catmap.maptree.add(
			&mapTreeNode{mapunit: mapunit},
		)
	}
	return &catmap
}

func (c *categoryMap) find(src int) int {
	dst := c.maptree.find(src)
	if dst >= 0 {
		return dst
	}
	// Business rule: if there's no mapping, dst = src
	return src
}

type mapTreeNode struct {
	mapunit *mapping
	left    *mapTreeNode
	right   *mapTreeNode
}

func (tree *mapTreeNode) add(node *mapTreeNode) {
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
func (tree *mapTreeNode) find(src int) int {
	if tree.mapunit.containsSource(src) {
		delta := src - tree.mapunit.sourceStart
		return tree.mapunit.destStart + delta
	}
	if tree.left != nil && src < tree.mapunit.sourceStart {
		return tree.left.find(src)
	}
	if tree.right != nil && src > tree.mapunit.sourceStart {
		return tree.right.find(src)
	}
	return -1
}
