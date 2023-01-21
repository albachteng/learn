package tree

import (
	"errors"
	"fmt"
	"strings"
)

var ErrorNotFound = errors.New("could not find")

type TreeNode[T comparable] struct {
	val      T
	children []TreeNode[T]
}

func NewTreeNode[T comparable](val T) TreeNode[T] {
  n := new(TreeNode[T])
  n.val = val
  return *n
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
	return ErrorNotFound, nil
}

// removes a node and elevates its children, if any, in its place. Will not remove the root node
func (n *TreeNode[T]) RemoveNode(val T) error {
	for i, child := range n.children {
		if err := child.RemoveNode(val); child.val == val {
			if err != nil {
				updatedChildren := append(n.children[:i], child.children...)
				updatedChildren = append(updatedChildren, n.children[i+1:]...)
				n.children = updatedChildren
				return nil
			}
			return err
		}
	}
	return ErrorNotFound
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

type indentCount struct {
	count int
}

func (n *TreeNode[T]) printWithCount(count int) {
  indent := strings.Repeat(" ", count*2)
  fmt.Println(indent, n.val)
  for _, child := range n.children {
    child.printWithCount(count + 1)
  }
}

func (n *TreeNode[T]) Print() {
  n.printWithCount(0)
}

