package pinkgray

import (
	"fmt"
	"testing"
)

type it int

func (i it) Less(j Item) bool {
	return i < j.(it)
}

func Test_simple(t *testing.T) {
	tr := Tree{}
	tr.Insert(it(10))
	tr.Insert(it(12))
	tr.Insert(it(11))
	fmt.Printf("root %v isBlack: %t\n", tr.root.value, tr.root.isBlack)
	fmt.Printf("left %v isBlack: %t\n", tr.root.left.value, tr.root.left.isBlack)
	fmt.Printf("right %v isBlack: %t\n", tr.root.right.value, tr.root.right.isBlack)
}
