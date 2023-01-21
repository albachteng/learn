package tree

type BST[T comparable] struct {
	val         T
	left, right *BST[T]
}

func NewBST[T comparable]() BST[T] {
	return *new(BST[T])
}

func (b *BST[T]) SetVal(val T) {
	b.val = val
}

func (b *BST[T]) GetVal() T {
	return b.val
}

func Sort[T comparable](values []T) {
	var root *BST[T]
	for _, v := range values {
		root = root.add(v)
	}
	root.appendValues(values[:0], root)
}

func (b *BST[T]) appendValues(values []T, t *BST[T]) []T {
	if b != nil {
		values = b.appendValues(values, b.left)
		values = append(values, b.val)
		values = b.appendValues(values, b.right)
	}
	return values
}

func (b *BST[T]) add(val T) *BST[T] {
	if val < b.val {
		b.left = b.left.add(val)
	} else {
		b.right = b.right.add(val)
	}
	return b
}
