package btree

import "sort"

type node struct {
	keys     []int
	children []*node
}

type BTree struct {
	root  *node
	order int
}

func (b *BTree) Insert(key int) {
	if b.root == nil {
		b.root = &node{keys: []int{key}}
		return
	}
	b.root = b.root.insert(key, b.order)
}

func (n *node) insert(key int, order int) *node {
	i := sort.Search(len(n.keys), func(i int) bool {
		return n.keys[i] >= key
	})
	// i is within the range of n.keys and the key is already present
	if i < len(n.keys) && n.keys[i] == key {
		// the current node is the correct node and no insertion is necessary, return as-is
		return n
	}
	// this node has no children and either the index is out of range or the key was not already present
	if len(n.children) == 0 {
		// add a zero to the end of this node's keys
		n.keys = append(n.keys, 0)
		// shift up by one
		copy(n.keys[i+1:], n.keys[i:])
		// add the index to the end, that is the right place
		n.keys[i] = key
		// if we've exceeded the order we must split
		if len(n.keys) > order {
			return n.split(order)
		}
		// after we've split, return the now corrected node
		return n
	}
	// if this node has children and the index was out of range or not already present, we add to children
	// I am unsure why we are using the index we found in the sorted search above here, I thought that was the index for the keys...
	n.children[i] = n.children[i].insert(key, order)
	// if this takes the children past the order threshold, we split the child node at i
	// again I am unsure why we are using the i we found in the sorted search above
	if len(n.children[i].keys) > order {
		return n.splitChild(i, order)
	}
	return n
}

func (n *node) split(order int) *node {
	mid := len(n.keys) / 2
	parent := &node{
		keys: []int{n.keys[mid]}, // parent node's keys array is initialized to only the midpoint key
		children: []*node{
			n, // its children are the current node...
			{ // ... plus a new node...
				keys:     n.keys[mid+1:], // whose keys are the second half of the current node's keys
				children: nil,            // and which has no children
			},
		},
	}
	n.keys = n.keys[:mid]           // the current node's keys should now not include the second half
	n.children = n.children[:mid+1] // the current node's children should also include only the second half
	// the parent might need to be split as well
	if len(parent.keys) > order {
		return parent.split(order)
	}
	return parent
}

func (n *node) splitChild(i, order int) *node {
	child := n.children[i]
	mid := len(child.keys) / 2
	sibling := &node{
		keys:     child.keys[mid+1:],     // sibling has the second half of the child keys
		children: child.children[mid+1:], // and second half of the child children
	}
	child.keys = child.keys[:mid]           // update child keys to exclude the sibling's
	child.children = child.children[:mid+1] // update child children to exclude the sibling's
	n.keys = append(n.keys, 0)              // add a nil value key to the end of n's keys
	copy(n.keys[i+1:], n.keys[i:])          // shift up by one, creating space for the new item
	copy(n.children[i+2:], n.children[i+1:])
	n.children[i+1] = sibling                 // insert the new sibling
	n.keys[i] = child.keys[len(child.keys)-1] // not sure
	if len(n.keys) > order {
		return n.split(order)
	}
	return n
}

func (b *BTree) Search(key int) bool {
	if b.root == nil {
		return false
	}
	return b.root.search(key)
}

func (n *node) search(key int) bool {
	i := sort.Search(len(n.keys), func(i int) bool {
		return n.keys[i] >= key
	})
	if i < len(n.keys) && n.keys[i] == key {
		return true
	}
	if len(n.children) == 0 {
		return false
	}
	return n.children[i].search(key)
}
