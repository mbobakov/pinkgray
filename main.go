// Package pinkgray is Red-Black tree implementation.
// It is very usefull cause can operate any data due implementing Item
// interface
// It repeats beutify google/btree API
// If i.Less(j) == j.Less(i) is the equal nodes
package pinkgray

type Tree struct {
	root *node
	size int
}

// Item for the tree
type Item interface {
	Less(Item) bool
}

type node struct {
	isBlack, isLeft, isRight bool
	parent                   *node
	left                     *node
	right                    *node
	value                    Item
}

func (t *Tree) Insert(i Item) {
	if t.root == nil {
		t.root = &node{isBlack: true, value: i}
		return
	}
	p := t.root.findParentFor(i)

	//TODO: equal case
	if p.value.Less(i) {
		p.right = &node{parent: p, isRight: true, value: i}
		t.fix(p.right)
		return
	}
	p.left = &node{parent: p, isLeft: true, value: i}
	t.fix(p.left)
}

func (n *node) findParentFor(i Item) *node {
	if n.value.Less(i) {
		if n.right == nil { // terminate recursion
			return n
		}
		return n.right.findParentFor(i)
	}
	if n.left == nil { // terminate recursion
		return n
	}
	return n.left.findParentFor(i)
}

// Fix always must called for the cause node
func (t *Tree) fix(n *node) {
	aunt := n.aunt()
	if aunt == nil {
		t.rotate(n)
		return
	}

}

func (n *node) colorFlip() {

}

func (t *Tree) rotate(n *node) {
	if n.grandParent() == nil {
		return
	}
	if n.isLeft && n.parent.isRight {
		// right rotate
		n.parent.parent.right = n
		n.isLeft = false
		n.isRight = true
		n.right = n.parent
		n.parent.left = nil
		n.parent = n.parent.parent
		// left rotate
		n.left = n.parent
		n.parent = n.parent.parent
		n.left.parent = n
		n.left.isLeft = true
		n.left.isRight = false
		n.isBlack = true
		n.left.isBlack = false
		n.right.isBlack = false
	}
	if n.parent == nil { //isRoot
		t.root = n
	}

}

func (n *node) aunt() *node {
	gp := n.grandParent()
	if gp == nil {
		return nil
	}
	if n.parent.isLeft {
		return gp.right
	}
	return gp.left
}

func (n *node) grandParent() *node {
	if n.parent == nil || n.parent.parent == nil {
		return nil
	}
	return n.parent.parent
}
