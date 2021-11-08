package tree

import "testing"

func TestNewTree(t *testing.T) {
	tree := NewTree(&Node{Value: 1})
	for i := 0; i < 10; i += 2 {
		tree.InsertNode(int64(i))
	}
	tree.Print()
	searchNode := tree.Search(4)
	println(searchNode.Value)
	searchNode.Insert(3)
	tree.Print()
}
