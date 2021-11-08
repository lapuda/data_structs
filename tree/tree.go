// Package tree
// 二叉树的实现
// 1) 若当前的二叉查找树为空，则插入的元素为根节点;
// 2) 若插入的元素值小于根节点值，则将元素插入到左子树中;
// 3) 若插入的元素值不小于根节点值，则将元素插入到右子树中。
package tree

import "fmt"

type Node struct {
	Value    int64
	Left     *Node
	Right    *Node
	IsBranch bool
	IsLeaf   bool
	IsRoot   bool
	Dep      int
}

// Insert 插入节点
func (thisNode *Node) Insert(insertValue int64) {
	insertNode := Node{Value: insertValue, IsBranch: false, IsLeaf: true}
	// 相同值的节点跳过
	if thisNode.Value == insertValue {
		return
	}
	// 如果插入节点小于当前节点
	if thisNode.More(insertNode) {
		if thisNode.Left == nil {
			thisNode.Left = &insertNode
			thisNode.Dep++
		} else {
			thisNode.Left.IsBranch = true
			thisNode.Left.IsLeaf = false
			thisNode.Left.Insert(insertValue)
		}
	}
	// 如果插入节点大于当前节点
	if thisNode.Less(insertNode) {
		if thisNode.Right == nil {
			thisNode.Right = &insertNode
			thisNode.Dep++
		} else {
			thisNode.Right.IsBranch = true
			thisNode.Right.IsLeaf = false
			thisNode.Right.Insert(insertValue)
		}
	}
}

// Less 是否比传入节点小
func (thisNode *Node) Less(comparedNode Node) bool {
	return thisNode.Value < comparedNode.Value
}

// More 是否比传入节点大
func (thisNode *Node) More(comparedNode Node) bool {
	return thisNode.Value > comparedNode.Value
}

// Search 查找
func (thisNode *Node) Search(searchValue int64) *Node {
	if thisNode.Value == searchValue {
		return thisNode
	}
	if searchValue > thisNode.Value {
		return thisNode.Right.Search(searchValue)
	}
	return thisNode.Left.Search(searchValue)
}

func (thisNode *Node) Print() {
	print(fmt.Sprintf("value:%v,is_root:%v,is_leaf:%v,is_branch:%v \n",
		thisNode.Value, thisNode.IsRoot, thisNode.IsLeaf, thisNode.IsBranch))
	if thisNode.Left != nil {
		thisNode.Left.Print()
	}
	if thisNode.Right != nil {
		thisNode.Right.Print()
	}
}

type Tree struct {
	Root    *Node
	LeafNum int64
}

// InsertNode 插入节点
func (thisTree *Tree) InsertNode(insertValue int64) {
	thisTree.Root.Insert(insertValue)
}

// Search 对数据进行搜索
func (thisTree *Tree) Search(searchValue int64) *Node {
	return thisTree.Root.Search(searchValue)
}

func (thisTree *Tree) Print() {
	thisTree.Root.Print()
}

func NewTree(node *Node) *Tree {
	node.IsRoot = true
	return &Tree{Root: node}
}
