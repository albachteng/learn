package tree

import (
	"errors"
)

var ErrorNotFound = errors.New("could not find")

type TreeNode[T comparable] struct {
	val      T
	children []TreeNode[T]
}

func NewTreeNode[T comparable](val T) TreeNode[T] {
	return TreeNode[T]{val: val, children: make([]TreeNode[T], 0)}
}

func (n *TreeNode[T]) GetValue() T {
	return n.val
}

func (n *TreeNode[T]) SetValue(val T) {
	n.val = val
}

func (n *TreeNode[T]) AddChild(val T) {
	n.children = append(n.children, NewTreeNode(val))
}

func (n *TreeNode[T]) Contains(val T) bool {
	if n.val == val {
		return true
	}
	for _, child := range n.children {
		if child.Contains(val) {
			return true
		}
	}
	return false
}

func (n *TreeNode[T]) GetNode(val T) (error, *TreeNode[T]) {
	if n.val == val {
		return nil, n
	}
	for _, child := range n.children {
		err, found := child.GetNode(val)
		if err == nil {
			return nil, found
		}
	}
	return errors.New("could not find"), nil
}

// removes a node and elevates its children, if any, in its place will not remove the root node
func (n *TreeNode[T]) RemoveNode(val T) error {
	for i, child := range n.children {
		if err := child.RemoveNode(val); child.val == val {
			if err != nil {
				updatedChildren := n.children[:i]
				updatedChildren = append(updatedChildren, child.children...)
				updatedChildren = append(updatedChildren, n.children[i+1:]...)
				n.children = updatedChildren
				return nil
			}
			return err
		}
	}
	return errors.New("could not find")
}

func (n *TreeNode[T]) RemoveBranchByNode(val T) error {
	for i, child := range n.children {
		if err := child.RemoveBranchByNode(val); child.val == val {
			if err != nil {
				n.children = append(n.children[:i], n.children[i+1:]...)
				return nil
			}
		}
	}
	return ErrorNotFound
}
